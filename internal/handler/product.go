package handler

import (
	"context"
	"database/sql"
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
	UpdateProduct() error
	DeleteProduct() error
	GetProductList(ctx context.Context) ([]*domain.Product, error)
	GetProductDetail() error
}

func NewProductHandler(repo ProductRepository) ProductHandler {
	return ProductHandler{repo: repo}
}

func (ph *ProductHandler) Home(c *gin.Context) {
	ctx := context.Background()
	products, err := ph.repo.GetProductList(ctx)
	if err != nil {
		//todo
	}
	_ = products

	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "상품 리스트",
		//""
	})
	return
}

func (ph *ProductHandler) CreateProductPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product_create.tmpl", gin.H{
		"title": "메뉴 추가하기",
	})
	return
}
func (ph *ProductHandler) CreateProduct(c *gin.Context) {
	ctx := context.Background()
	category := c.PostForm("category")
	price := c.PostForm("price")
	cost := c.PostForm("cost")
	name := c.PostForm("name")
	description := c.PostForm("description")
	barcode := c.PostForm("barcode")
	expiredAt := c.PostForm("expired_at")
	_ = expiredAt
	size := c.PostForm("size")

	p, _ := strconv.Atoi(price)
	co, _ := strconv.Atoi(cost)

	product := &domain.Product{
		Category:    category,
		Price:       int64(p),
		Cost:        int64(co),
		Name:        name,
		Description: description,
		Barcode:     barcode,
		ExpiredAt:   sql.NullTime{Time: time.Now(), Valid: true},
		Size:        domain.ProductSize(size),
	}

	if err := ph.repo.CreateProduct(ctx, product); err != nil {
		//todo
	}
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"title": "Main website",
	})
	return
}

func (ph *ProductHandler) Update(c *gin.Context) {
	return
}

func (ph *ProductHandler) Delete(c *gin.Context) {
	return
}

func (ph *ProductHandler) Search(c *gin.Context) {
	return
}
