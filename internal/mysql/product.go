package mysql

import (
	"context"
	"payhere/internal/domain"
)

func (r ProductRepo) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}

func (r ProductRepo) UpdateProduct() error {
	return nil
}

func (r ProductRepo) DeleteProduct() error {
	return nil
}

func (r ProductRepo) GetProductList(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.DB.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r ProductRepo) GetProductDetail() error {
	return nil
}
