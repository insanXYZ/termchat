package handler

import (
	"bin-term-chat/model"
	"bytes"
	"encoding/json"
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

func (h *Handler) UpdateUser(req *model.UpdateUser, token string) (*model.Response, error) {
	if req.Name == "" && req.Email == "" && req.Password == "" && req.Bio == "" {
		return nil, nil
	}

	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return h.NewRequest(http.MethodPut, h.Url+"/api/user", bytes.NewReader(marshal), func(request *http.Request) {
		request.Header.Set("Authorization", "Bearer "+token)
	})
}
