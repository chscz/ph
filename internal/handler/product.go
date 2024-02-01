package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"payhere/internal/domain"
	"strconv"
	"time"
)

type ProductHandler struct {
	repo ProductRepository
}

type ProductRepository interface {
	CreateProduct(ctx context.Context, product *domain.Product) error
	UpdateProduct(ctx context.Context, product *domain.Product) error
	DeleteProduct(ctx context.Context, id int) error
	GetProduct(ctx context.Context, id int) (domain.Product, error)
	GetProductAllList(ctx context.Context) ([]*domain.Product, error)
	GetProductSearchList(ctx context.Context, keyword string) ([]*domain.Product, error)
	GetProductSearchListByChoSung(ctx context.Context, keyword string) ([]*domain.Product, error)
}

func NewProductHandler(repo ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

func (ph *ProductHandler) Home(c *gin.Context) {
	// 여기

	ctx := context.Background()
	products, err := ph.repo.GetProductAllList(ctx)
	if err != nil {
		//todo
	}

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title":    "상품 리스트",
		"products": convertFromDomainProductList(products),
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
	ctx := context.Background()
	keyword := c.Query("search_by_name")

	var products []*domain.Product
	var err error
	if isOnlyChoSung(keyword) {
		products, err = ph.repo.GetProductSearchListByChoSung(ctx, keyword)
		if err != nil {
			//todo
		}
	} else {
		products, err = ph.repo.GetProductSearchList(ctx, keyword)
		if err != nil {
			//todo
		}
	}

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title":         "상품 검색 결과",
		"SearchKeyword": keyword,
		"products":      convertFromDomainProductList(products),
	})
}
