package mysql

import (
	"context"

	"github.com/chscz/ph/internal/domain"
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

func (r ProductRepo) GetProducts(ctx context.Context, itemsPerPage int, whereClause string) ([]*domain.Product, error) {
	var products []*domain.Product

	query := r.DB.WithContext(ctx).
		Where(whereClause).
		Order("created_at DESC, id DESC")

	if itemsPerPage > 0 {
		query = query.Limit(itemsPerPage)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductRepo) GetTotalProductCount(ctx context.Context) (int, error) {
	var cnt int64
	if err := r.DB.WithContext(ctx).
		Model(domain.Product{}).
		Count(&cnt).
		Error; err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (r ProductRepo) GetProductSearchList(ctx context.Context, keyword string, itemsPerPage int, whereClause string) ([]*domain.Product, error) {
	var products []*domain.Product

	query := r.DB.WithContext(ctx).
		Where("name LIKE ?", "%"+keyword+"%").
		Where(whereClause).
		Order("created_at DESC, id DESC")

	if itemsPerPage > 0 {
		query = query.Limit(itemsPerPage)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductRepo) GetTotalSearchedProductsCount(ctx context.Context, keyword string) (int, error) {
	var cnt int64
	if err := r.DB.WithContext(ctx).
		Model(domain.Product{}).
		Where("name LIKE ?", "%"+keyword+"%").
		Count(&cnt).
		Error; err != nil {
		return 0, err
	}
	return int(cnt), nil
}

func (r ProductRepo) GetProductSearchListByChoSung(ctx context.Context, keyword string, itemsPerPage int, whereClause string) ([]*domain.Product, error) {
	var products []*domain.Product

	query := r.DB.WithContext(ctx).
		Where("FNC_CHOSUNG(name) LIKE ?", "%"+keyword+"%").
		Where(whereClause).
		Order("created_at DESC, id DESC")

	if itemsPerPage > 0 {
		query = query.Limit(itemsPerPage)
	}

	if err := query.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r ProductRepo) GetTotalSearchedProductsCountByChoSung(ctx context.Context, keyword string) (int, error) {
	var cnt int64
	if err := r.DB.WithContext(ctx).
		Model(domain.Product{}).
		Where("FNC_CHOSUNG(name) LIKE ?", "%"+keyword+"%").
		Count(&cnt).
		Error; err != nil {
		return 0, err
	}
	return int(cnt), nil
}
