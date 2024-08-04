package handler

import (
	"bin-term-chat/model"
	"bytes"
	"encoding/json"
	"net/http"
)

func (h *Handler) GetUserWithId(input, token string) (*model.Response, error) {

	var query string

	if string(input[0]) == "#" {
		query = "?id=" + string(input[1:])
	} else {
		query = "?name=" + input
	}

	return h.NewRequest(http.MethodGet, h.Url+"/api/user"+query, nil, func(request *http.Request) {
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
