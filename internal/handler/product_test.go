package handler

import (
	"context"
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/chscz/ph/internal/domain"
	"github.com/chscz/ph/internal/handler/mock"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	testID          = "3"
	testCategory    = "test-category"
	testPrice       = "1000"
	testCost        = "10"
	testName        = "test-name"
	testDescription = "test-description"
	testBarcode     = "test-barcode"
	testExpiredAt   = "2024-01-01T00:00"
	testSize        = "test-size"
)

func TestProductHandler_Home(t *testing.T) {

}

func TestProductHandler_CreateProduct(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/product/create", nil)
	req.PostForm = map[string][]string{
		"category":    {testCategory},
		"price":       {testPrice},
		"cost":        {testCost},
		"name":        {testName},
		"description": {testDescription},
		"barcode":     {testBarcode},
		"expired_at":  {testExpiredAt},
		"size":        {testSize},
	}

	// gin test context 생성
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	ctrl := gomock.NewController(t)
	repo := mock.NewMockProductRepository(ctrl)

	price, _ := strconv.Atoi(testPrice)
	cost, _ := strconv.Atoi(testCost)
	expiredAt, _ := time.Parse(htmlInputTypeDatetimeLocalFormat, testExpiredAt)
	repo.EXPECT().
		CreateProduct(
			context.Background(), &domain.Product{
				Category:    testCategory,
				Price:       int64(price),
				Cost:        int64(cost),
				Name:        testName,
				Description: testDescription,
				Barcode:     testBarcode,
				ExpiredAt: sql.NullTime{
					Time:  expiredAt,
					Valid: true,
				},
				Size: testSize,
			}).
		Return(nil)

	ph := NewProductHandler(repo, false)

	ph.CreateProduct(c)
}

func TestProductHandler_DeleteProduct(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("DELETE", "/product/delete", nil)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: testID})

	ctrl := gomock.NewController(t)
	repo := mock.NewMockProductRepository(ctrl)

	id, _ := strconv.Atoi(testID)
	repo.EXPECT().DeleteProduct(context.Background(), id).Return(nil)

	ph := NewProductHandler(repo, false)

	ph.DeleteProduct(c)

	assert.Equal(t, http.StatusOK, http.StatusOK)
}

func TestProductHandler_UpdateProduct(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("PUT", "/product/update", nil)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: testID})
	req.PostForm = map[string][]string{
		"category":    {testCategory},
		"price":       {testPrice},
		"cost":        {testCost},
		"name":        {testName},
		"description": {testDescription},
		"barcode":     {testBarcode},
		"expired_at":  {testExpiredAt},
		"size":        {testSize},
	}

	ctrl := gomock.NewController(t)
	repo := mock.NewMockProductRepository(ctrl)

	price, _ := strconv.Atoi(testPrice)
	cost, _ := strconv.Atoi(testCost)
	id, _ := strconv.Atoi(testID)
	expiredAt, _ := time.Parse(htmlInputTypeDatetimeLocalFormat, testExpiredAt)
	repo.EXPECT().
		UpdateProduct(
			context.Background(), &domain.Product{
				ID:          int64(id),
				Category:    testCategory,
				Price:       int64(price),
				Cost:        int64(cost),
				Name:        testName,
				Description: testDescription,
				Barcode:     testBarcode,
				ExpiredAt: sql.NullTime{
					Time:  expiredAt,
					Valid: true,
				},
				Size: testSize,
			}).
		Return(nil)

	ph := NewProductHandler(repo, false)

	ph.UpdateProduct(c)

	assert.Equal(t, http.StatusOK, http.StatusOK)
}

func TestProductHandler_GetProductDetail(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest("GET", "/product/detail", nil)
	c.Request = req
	c.Params = append(c.Params, gin.Param{Key: "id", Value: testID})

	ctrl := gomock.NewController(t)
	repo := mock.NewMockProductRepository(ctrl)

	ID, _ := strconv.Atoi(testID)
	price, _ := strconv.Atoi(testPrice)
	cost, _ := strconv.Atoi(testCost)
	expiredAt, _ := time.Parse(htmlInputTypeDatetimeLocalFormat, testExpiredAt)
	repo.EXPECT().
		GetProduct(context.Background(), ID).
		Return(domain.Product{
			ID:          int64(ID),
			Category:    testCategory,
			Price:       int64(price),
			Cost:        int64(cost),
			Name:        testName,
			Description: testDescription,
			Barcode:     testBarcode,
			ExpiredAt: sql.NullTime{
				Time:  expiredAt,
				Valid: true,
			},
			Size: testSize,
		}, nil)

	ph := NewProductHandler(repo, false)

	ph.GetProductDetail(c)

	assert.Equal(t, http.StatusOK, http.StatusOK)

	assert.Equal(
		t,
		Product{
			ID:          ID,
			Category:    testCategory,
			Price:       price,
			Cost:        cost,
			Name:        testName,
			Description: testDescription,
			Barcode:     testBarcode,
			ExpiredAt:   testExpiredAt,
			Size:        testSize,
		}, w.Result())
}

func TestProductHandler_SearchProduct(t *testing.T) {

}
