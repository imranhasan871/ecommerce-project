package user

type Handler struct{}

func NewHandler() *Handler {
	handler := &Handler{}
	return handler
}
