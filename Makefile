mocks:
	mockgen --source=internal/handler/user.go --destination=internal/handler/mock/user.go -package=mock
	mockgen --source=internal/handler/product.go --destination=internal/handler/mock/product.go -package=mock
