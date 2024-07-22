package handler

import (
	"bin-term-chat/model"
	"errors"
	"net/http"
)

func (h *Handler) GetUserWithId(id, token string) (*model.Response, error) {

	if id == "" {
		return nil, errors.New("id required")
	}

	return h.NewRequest(http.MethodGet, h.Url+"/api/user?id="+id, nil, func(request *http.Request) {
		request.Header.Set("Authorization", "Bearer "+token)
	})
}
