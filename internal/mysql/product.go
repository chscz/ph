package mysql

import (
	"context"
	"payhere/internal/domain"
)

func (r ProductRepo) CreateProduct(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Create(product).Error
}

func (r ProductRepo) UpdateProduct(ctx context.Context, product *domain.Product) error {
	return r.DB.WithContext(ctx).Updates(product).Error
}

func (r ProductRepo) DeleteProduct(ctx context.Context, id int) error {
	return r.DB.WithContext(ctx).Model(domain.Product{}).Delete("id = ?", id).Error
}

func (r ProductRepo) GetProductAllList(ctx context.Context) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.DB.WithContext(ctx).
		Order("created_at DESC").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r ProductRepo) GetProduct(ctx context.Context, id int) (domain.Product, error) {
	var product domain.Product
	if err := r.DB.WithContext(ctx).
		Where("id = ?", id).
		Take(&product).
		Error; err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r ProductRepo) GetProductSearchList(ctx context.Context, keyword string) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.DB.WithContext(ctx).
		Model(domain.Product{}).
		Where("name LIKE ?", "%"+keyword+"%").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r ProductRepo) GetProductSearchListByChoSung(ctx context.Context, keyword string) ([]*domain.Product, error) {
	var products []*domain.Product
	if err := r.DB.WithContext(ctx).
		Model(domain.Product{}).
		Where("FNC_CHOSUNG(name) LIKE ?", "%"+keyword+"%").
		Find(&products).
		Error; err != nil {
		return nil, err
	}
	return products, nil
}
