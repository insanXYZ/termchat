package handler

import (
	"bin-term-chat/model"
	"net/http"
)

func (h *Handler) GetChats(token string) (*model.Response, error) {
	return h.NewRequest(http.MethodGet, h.Url+"/api/chat", nil, func(request *http.Request) {
		request.Header.Set("Authorization", "Bearer "+token)
	})
}
