package handler

type MenuHandler struct {
	repo MenuRepository
}

type MenuRepository interface {
}

func NewMenuHandler(repo MenuRepository) MenuHandler {
	return MenuHandler{repo: repo}
}
