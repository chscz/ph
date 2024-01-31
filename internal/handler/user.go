package handler

type UserHandler struct {
	repo UserRepository
}

type UserRepository interface {
}

func NewUserHandler(repo UserRepository) UserHandler {
	return UserHandler{repo: repo}
}
