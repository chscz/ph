package handler

import (
	"database/sql"
	"payhere/internal/domain"
	"strings"
	"time"
	"unicode"
)

const htmlInputTypeDatetimeLocalFormat = "2006-01-02T15:04"

var hangulChoSung = "ㄱㄲㄴㄷㄸㄹㅁㅂㅃㅅㅆㅇㅈㅉㅊㅋㅌㅍㅎ"

type ProductSize string

const (
	ProductSizeLarge ProductSize = "large"
	ProductSizeSmall ProductSize = "small"
)

type Product struct {
	ID          int
	Category    string
	Price       int
	Cost        int
	Name        string
	Description string
	Barcode     string
	ExpiredAt   string
	Size        ProductSize
}

func (p Product) convertToDomainModel() (*domain.Product, error) {
	t, err := time.Parse(htmlInputTypeDatetimeLocalFormat, p.ExpiredAt)
	if err != nil {
		return &domain.Product{}, err
	}
	return &domain.Product{
		ID:          p.ID,
		Category:    p.Category,
		Price:       int64(p.Price),
		Cost:        int64(p.Cost),
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		ExpiredAt: sql.NullTime{
			Time:  t,
			Valid: true,
		},
		Size: string(p.Size),
	}, nil
}

func convertFromDomainProductList(products []*domain.Product) []*Product {
	productList := make([]*Product, len(products))
	for i, product := range products {
		productList[i] = convertFromDomainProduct(product)
	}
	return productList
}

func convertFromDomainProduct(p *domain.Product) *Product {
	return &Product{
		ID:          p.ID,
		Category:    p.Category,
		Price:       int(p.Price),
		Cost:        int(p.Cost),
		Name:        p.Name,
		Description: p.Description,
		Barcode:     p.Barcode,
		ExpiredAt:   p.ExpiredAt.Time.Format(htmlInputTypeDatetimeLocalFormat),
		Size:        ProductSize(p.Size),
	}
}

func isOnlyChoSung(str string) bool {
	for _, char := range str {
		if !unicode.Is(unicode.Hangul, char) {
			return false
		}
		if !strings.Contains(hangulChoSung, string(char)) {
			return false
		}
	}
	return true
}
