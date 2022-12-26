package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stay/core"
	"stay/internal/ports"
)

type propertyRequest struct {
	ID          string
	UserUID     string `json:"user_uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Typ         string `json:"type"`
	Rooms       int    `json:"rooms"`
}

func NewPropertyCreateRequest(c echo.Context) (*propertyRequest, *core.Error) {
	var req *propertyRequest

	if err := c.Bind(&req); err != nil {
		return nil, core.NewError(errBadPropertyRequest, http.StatusBadRequest)
	}

	return req, nil
}

func (r *propertyRequest) ToDomain() *ports.Property {
	return &ports.Property{
		ID:          r.ID,
		UserUID:     r.UserUID,
		Name:        r.Name,
		Description: r.Description,
		Address:     r.Address,
		Typ:         r.Typ,
		Rooms:       r.Rooms,
	}
}
