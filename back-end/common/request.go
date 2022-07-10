package common

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"io/ioutil"
)

type APIRequest interface {
	GetContent(data interface{}) error
	GetContentText() string
}

type HTTPAPIRequest struct {
	context echo.Context
	body    string
}

func (req *HTTPAPIRequest) GetContent(data interface{}) error {
	return json.Unmarshal([]byte(req.GetContentText()), data)
}

func (req *HTTPAPIRequest) GetContentText() string {
	var bodyBytes []byte
	reqBody := req.context.Request().Body
	if reqBody != nil {
		bodyBytes, _ = ioutil.ReadAll(reqBody)
	}

	req.body = string(bodyBytes)

	return req.body
}
