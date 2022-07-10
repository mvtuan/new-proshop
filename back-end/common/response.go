package common

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type APIResponder interface {
	Respond()
}

type HTTPApiResponder struct {
	context echo.Context
}

func (res HTTPApiResponder) Respond(response *APIResponse) error {
	var context = res.context
	switch response.Status {
	case APIStatus.Ok:
		return context.JSON(http.StatusOK, response)
	case APIStatus.Error:
		return context.JSON(http.StatusInternalServerError, response)
	case APIStatus.NotFound:
		return context.JSON(http.StatusNotFound, response)
	case APIStatus.Invalid:
		return context.JSON(http.StatusBadRequest, response)
	case APIStatus.Forbidden:
		return context.JSON(http.StatusForbidden, response)
	case APIStatus.Unauthorized:
		return context.JSON(http.StatusUnauthorized, response)
	}

	return context.JSON(http.StatusBadRequest, response)
}
