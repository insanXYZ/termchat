package handler

import (
	"bin-term-chat/model"
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

func (h *Handler) Login(req *model.ReqLogin) (*model.Response, error) {
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpresp, err := http.Post(h.Url+"/api/login", "application/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}

	resp := new(model.Response)

	err = json.NewDecoder(httpresp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	if httpresp.StatusCode > 399 {
		return nil, errors.New(resp.Message)
	}

	return resp, nil

}

func (h *Handler) Register(req *model.ReqRegister) (*model.Response, error) {
	marshal, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	httpresp, err := http.Post(h.Url+"/api/register", "application/json", bytes.NewReader(marshal))
	if err != nil {
		return nil, err
	}

	resp := new(model.Response)

	err = json.NewDecoder(httpresp.Body).Decode(resp)
	if err != nil {
		return nil, err
	}

	if httpresp.StatusCode > 399 {
		return nil, errors.New(resp.Message)
	}

	return resp, nil

}
