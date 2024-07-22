package handler

import (
	"bin-term-chat/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Handler struct {
	Url string
}

func NewHandler(url string) *Handler {
	return &Handler{Url: url}
}

func (h *Handler) NewRequest(method, url string, body io.Reader, setRequest func(*http.Request)) (*model.Response, error) {
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	setRequest(request)

	client := http.Client{}
	httpres, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpres.Body.Close()

	response := new(model.Response)

	err = json.NewDecoder(httpres.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if httpres.StatusCode > 399 {
		return nil, errors.New(response.Message)
	}

	return response, nil

}
