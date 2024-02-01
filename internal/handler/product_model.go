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

func getFirstLastProductID(items []*domain.Product) (first, last int) {
	first = items[0].ID
	last = items[len(items)-1].ID
	return
}

func setProductPage(page, totalCount int) (prevPage, nextPage int) {
	// 이전 페이지와 다음 페이지를 계산
	prevPage = page - 1
	nextPage = page + 1

	// 이전 페이지가 1보다 작으면 이전 페이지는 1로 설정
	if prevPage < 1 {
		prevPage = 1
	}

	// 다음 페이지가 마지막 페이지를 넘어가면 다음 페이지는 마지막 페이지로 설정
	if nextPage > (totalCount+itemsPerPage-1)/itemsPerPage {
		nextPage = (totalCount + itemsPerPage - 1) / itemsPerPage
	}
	return
}
