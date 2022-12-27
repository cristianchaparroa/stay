package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stay/core"
)

type propertyParamRequest struct {
	ID  string
	UID string
}

func NewPropertyGetByIDRequest(c echo.Context) (*propertyParamRequest, *core.Error) {
	var req *propertyParamRequest

	id := c.QueryParam(propertyIDParam.ToString())
	if id != "" {
		req.ID = id
	} else {
		return nil, core.NewError(errBadPropertyRequestIDMissing, http.StatusBadRequest)
	}

	uid := c.QueryParam(propertyUIDParam.ToString())
	if uid != "" {
		req.UID = uid
		return req, nil
	}
	return nil, core.NewError(errBadPropertyRequestUIDMissing, http.StatusBadRequest)
}

func NewPropertyGetAllRequest(c echo.Context) (*propertyParamRequest, *core.Error) {
	req := &propertyParamRequest{}
	uid := c.Param(propertyUIDParam.ToString())

	if uid != "" {
		req.UID = uid
		return req, nil
	}

	return nil, core.NewError(errBadPropertyRequestUIDMissing, http.StatusBadRequest)
}
