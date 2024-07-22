package handler

import (
	"bin-term-chat/model"
	"bytes"
	"encoding/json"
	"net/http"
)

func (h *Handler) Login(req *model.ReqLogin) (*model.Response, error) {
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return h.NewRequest(http.MethodPost, h.Url+"/api/login", bytes.NewReader(marshal), func(request *http.Request) {
		request.Header.Set("Content-Type", "application/json")
	})

}

func (h *Handler) Register(req *model.ReqRegister) (*model.Response, error) {
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return h.NewRequest(http.MethodPost, h.Url+"/api/register", bytes.NewReader(marshal), func(request *http.Request) {
		request.Header.Set("Content-Type", "application/json")
	})

}
