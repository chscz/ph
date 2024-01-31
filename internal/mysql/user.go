package mysql

import (
	"context"
	"payhere/internal/domain"
)

func (r UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return r.DB.WithContext(ctx).Create(user).Error
}

func (r UserRepo) GetUser(ctx context.Context, phoneNumber string) (domain.User, error) {
	var user domain.User
	if err := r.DB.WithContext(ctx).
		Model(domain.User{}).
		Where("phone_number = ?", phoneNumber).
		Take(&user).
		Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}
