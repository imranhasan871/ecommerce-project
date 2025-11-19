package user

import "ecommerce/repo"

type Handler struct {
	userRepo repo.UserRepo
}

func NewHandler(userRepo repo.UserRepo) *Handler {
	handler := &Handler{
		userRepo: userRepo,
	}
	return handler
}
