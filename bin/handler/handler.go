package handler

type Handler struct {
	Url string
}

func NewHandler(url string) *Handler {
	return &Handler{Url: url}
}
