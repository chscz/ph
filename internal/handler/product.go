package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/internal/domain"
	"strconv"
	"time"
)

const itemsPerPage = 10

const (
	pageModeHome = iota
	pageModeNext
	pageModePrev
)

type ProductHandler struct {
	repo ProductRepository
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id int) error

	GetProduct(ctx context.Context, id int) (domain.Product, error)
	GetProducts(ctx context.Context, itemsPerPage int, whereClause string) ([]*domain.Product, error)
	GetTotalProductCount(ctx context.Context) (int, error)

	GetProductSearchList(ctx context.Context, keyword string, itemsPerPage int, whereClause string) ([]*domain.Product, error)
	GetTotalSearchedProductsCount(ctx context.Context, keyword string) (int, error)

	GetProductSearchListByChoSung(ctx context.Context, keyword string, itemsPerPage int, whereClause string) ([]*domain.Product, error)
	GetTotalSearchedProductsCountByChoSung(ctx context.Context, keyword string) (int, error)
}

func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (ph *ProductHandler) Home(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	// 페이징처리 모드
	mode := pageModeHome
	if c.Query("mode") == "next" {
		mode = pageModeNext
	} else if c.Query("mode") == "prev" {
		mode = pageModePrev
	}

	ctx := context.Background()
	products := make([]*domain.Product, 0)

	switch mode {
	case pageModeHome:
		products, err = ph.repo.GetProducts(ctx, itemsPerPage, "")
	case pageModeNext:
		cursor, _ := strconv.Atoi(c.Query("cursor"))
		products, err = ph.repo.GetProducts(ctx, itemsPerPage, fmt.Sprintf("id < %d", cursor))
	case pageModePrev:
		cursor, _ := strconv.Atoi(c.Query("cursor"))
		products, err = ph.repo.GetProducts(ctx, 0, fmt.Sprintf("id > %d", cursor))
		products = products[len(products)-itemsPerPage:]
	}
	if err != nil {
		//todo
	}

	if len(products) == 0 {
		c.HTML(http.StatusOK, "home.tmpl", gin.H{
			"title": "상품 리스트",
		})
		return
	}
	firstItemID, lastItemID := getFirstLastProductID(products)

	// 페이징 정보 계산
	totalCount, err := ph.repo.GetTotalProductCount(ctx)
	if err != nil {
		// 에러 처리
		//c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
		//	"error": "상품 전체 수를 가져오는 중에 오류가 발생했습니다.",
		//})
		return
	}
	prevPage, nextPage := setProductPage(page, totalCount)

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title":       "상품 리스트",
		"products":    convertFromDomainProductList(products),
		"firstItemID": firstItemID,
		"lastItemID":  lastItemID,
		"currentPage": page,
		"prevPage":    prevPage,
		"nextPage":    nextPage,
		"totalPages":  (totalCount + itemsPerPage - 1) / itemsPerPage,
	})
}

func (ph *ProductHandler) CreateProductPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product_create.tmpl", gin.H{
		"title": "메뉴 추가하기",
		"now":   time.Now().Format(htmlInputTypeDatetimeLocalFormat),
	})
}

func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	ctx := context.Background()

	price, _ := strconv.Atoi(c.PostForm("price"))
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	p := &Product{
		ID:          0,
		Category:    c.PostForm("category"),
		Price:       price,
		Cost:        cost,
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Barcode:     c.PostForm("barcode"),
		ExpiredAt:   c.PostForm("expired_at"),
		Size:        ProductSize(c.PostForm("size")),
	}

	product, err := p.convertToDomainModel()
	if err != nil {
		//todo
	}

	if err := ph.repo.CreateProduct(ctx, product); err != nil {
		//todo
	}
	c.Redirect(http.StatusFound, "/")
}

func (ph *ProductHandler) UpdateProductPage(c *gin.Context) {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))
	p, err := ph.repo.GetProduct(ctx, id)
	if err != nil {
		//todo
	}

	c.HTML(http.StatusOK, "product_update.tmpl", gin.H{
		"title":   "메뉴 수정하기",
		"product": convertFromDomainProduct(&p),
	})
}

func (ph *ProductHandler) UpdateProduct(c *gin.Context) {
	ctx := context.Background()

	id, _ := strconv.Atoi(c.Param("id"))
	price, _ := strconv.Atoi(c.PostForm("price"))
	cost, _ := strconv.Atoi(c.PostForm("cost"))
	p := &Product{
		ID:          id,
		Category:    c.PostForm("category"),
		Price:       price,
		Cost:        cost,
		Name:        c.PostForm("name"),
		Description: c.PostForm("description"),
		Barcode:     c.PostForm("barcode"),
		ExpiredAt:   c.PostForm("expired_at"),
		Size:        ProductSize(c.PostForm("size")),
	}

	product, err := p.convertToDomainModel()
	if err != nil {
		//todo
	}

	if err := ph.repo.UpdateProduct(ctx, product); err != nil {
		//todo
	}
	c.Redirect(http.StatusFound, "/")
}

func (ph *ProductHandler) DeleteProduct(c *gin.Context) {
	defer c.Redirect(http.StatusFound, "/")
	ctx := context.Background()
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	if err := ph.repo.DeleteProduct(ctx, id); err != nil {
		//todo
	}
}

func (ph *ProductHandler) GetProductDetail(c *gin.Context) {
	ctx := context.Background()
	paramID := c.Param("id")
	id, _ := strconv.Atoi(paramID)
	p, err := ph.repo.GetProduct(ctx, id)
	if err != nil {
		//todo
	}

	c.HTML(http.StatusOK, "product_detail.tmpl", gin.H{
		"title":   "상품 상세보기",
		"product": convertFromDomainProduct(&p),
	})
}

func (ph *ProductHandler) SearchProduct(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	// 페이징처리 모드
	mode := pageModeHome
	if c.Query("mode") == "next" {
		mode = pageModeNext
	} else if c.Query("mode") == "prev" {
		mode = pageModePrev
	}

	ctx := context.Background()
	products := make([]*domain.Product, 0)
	keyword := c.Query("search_by_name")

	var totalCount int
	if isOnlyChoSung(keyword) {
		switch mode {
		case pageModeHome:
			products, err = ph.repo.GetProductSearchListByChoSung(ctx, keyword, itemsPerPage, "")
		case pageModeNext:
			cursor, _ := strconv.Atoi(c.Query("cursor"))
			products, err = ph.repo.GetProductSearchListByChoSung(ctx, keyword, itemsPerPage, fmt.Sprintf("id < %d", cursor))

		case pageModePrev:
			cursor, _ := strconv.Atoi(c.Query("cursor"))
			products, err = ph.repo.GetProductSearchListByChoSung(ctx, keyword, 0, fmt.Sprintf("id > %d", cursor))
			products = products[len(products)-itemsPerPage:]
		}
		if err != nil {
			//todo
		}
		totalCount, err = ph.repo.GetTotalSearchedProductsCountByChoSung(ctx, keyword)
		if err != nil {
			// 에러 처리
			c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "상품 전체 수를 가져오는 중에 오류가 발생했습니다.",
			})
			return
		}

	} else {
		switch mode {
		case pageModeHome:
			products, err = ph.repo.GetProductSearchList(ctx, keyword, itemsPerPage, "")
		case pageModeNext:
			cursor, _ := strconv.Atoi(c.Query("cursor"))
			products, err = ph.repo.GetProductSearchList(ctx, keyword, itemsPerPage, fmt.Sprintf("id < %d", cursor))

		case pageModePrev:
			cursor, _ := strconv.Atoi(c.Query("cursor"))
			products, err = ph.repo.GetProductSearchList(ctx, keyword, 0, fmt.Sprintf("id > %d", cursor))
			products = products[len(products)-itemsPerPage:]
		}
		if err != nil {
			//todo
		}
		totalCount, err = ph.repo.GetTotalSearchedProductsCount(ctx, keyword)
		if err != nil {
			// 에러 처리
			c.HTML(http.StatusInternalServerError, "error.tmpl", gin.H{
				"error": "상품 전체 수를 가져오는 중에 오류가 발생했습니다.",
			})
			return
		}
	}
	firstItemID, lastItemID := getFirstLastProductID(products)

	prevPage, nextPage := setProductPage(page, totalCount)

	c.HTML(http.StatusOK, "product_search.tmpl", gin.H{
		"title":         "상품 검색 결과",
		"SearchKeyword": keyword,
		"products":      convertFromDomainProductList(products),
		"firstItemID":   firstItemID,
		"lastItemID":    lastItemID,
		"currentPage":   page,
		"prevPage":      prevPage,
		"nextPage":      nextPage,
		"totalPages":    (totalCount + itemsPerPage - 1) / itemsPerPage,
	})
}
