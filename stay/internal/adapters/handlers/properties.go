package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"stay/core"
	"stay/core/server"
	"stay/internal/ports"
)

type Properties struct {
	useCase ports.PropertyUseCase
}

func NewProperties(c ports.PropertyUseCase) server.Handler {
	return &Properties{useCase: c}
}

// Create It will add a new property associated with a specific user
//	@Summary		creates a property
//	@Description	Create will add a new property associated to specific user
//	@ID				users-properties-create
//	@Accept			json
//	@Produce		json
//	@Router			/users/properties [post]
//	@Param			body	body		propertyRequest	true	"This is the property information to create"
//	@Success		201		{object}	ports.Property
//	@Failure		400		{object}	core.Error
//	@Failure		422		{object}	core.Error
func (h *Properties) Create(c echo.Context) error {
	req, errReq := NewPropertyCreateRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	p, err := h.useCase.Save(req.ToDomain())
	if err != nil {
		e := core.NewError(err.Error(), http.StatusUnprocessableEntity)
		return c.JSON(e.Status, e)
	}

	return c.JSON(http.StatusCreated, p)
}

// GetAll will retrieve all the properties associated to specific user
//	@Summary		gets all properties
//	@Description	It will retrieve all properties associated to specific user
//	@ID				users-properties-get-all
//	@Accept			json
//	@Produce		json
//	@Router			/users/{uid}/properties [get]
//	@Param			uid	path		string	true	"User UID"
//	@Success		201	{object}	ports.Property
//	@Failure		400	{object}	core.Error
//	@Failure		422	{object}	core.Error
func (h *Properties) GetAll(c echo.Context) error {
	req, errReq := NewPropertyGetAllRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	ps, err := h.useCase.GetAll(req.UID)

	if err != nil {
		e := core.NewError(err.Error(), http.StatusUnprocessableEntity)
		return c.JSON(e.Status, e)
	}

	response := NewPropertyListResponse(ps)
	return c.JSON(http.StatusOK, response)
}

// GetByID get a specific property associated to one user.
//	@Summary		gets a property by id
//	@Description	It gets a specific property associated to one user.
//	@ID				users-properties-get-by-id
//	@Accept			json
//	@Produce		json
//	@Router			/users/{uid}/properties/{id} [get]
//	@Param			uid	path		string	true	"User UID"
//	@Param			id	query		string	true	"Property ID"
//	@Success		201	{object}	ports.Property
//	@Failure		400	{object}	core.Error
//	@Failure		422	{object}	core.Error
func (h *Properties) GetByID(c echo.Context) error {
	req, errReq := NewPropertyGetByIDRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	p, err := h.useCase.GetProperty(req.UID, req.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err)
	}

	return c.JSON(http.StatusOK, NewPropertyResponse(p))
}

// DeleteByID delete a property.
//	@Summary		deletes a property
//	@Description	It deletes a property
//	@ID				users-properties-delete
//	@Accept			json
//	@Produce		json
//	@Router			/users/{uid}/properties/{id} [get]
//	@Param			uid	path		string	true	"User UID"
//	@Param			id	query		string	true	"Property ID"
//	@Success		200	{object}	ports.Property
//	@Failure		400	{object}	core.Error
//	@Failure		422	{object}	core.Error
func (h *Properties) DeleteByID(c echo.Context) error {
	req, errReq := NewPropertyGetByIDRequest(c)
	if errReq != nil {
		return c.JSON(errReq.Status, errReq)
	}

	err := h.useCase.Delete(req.UID, req.ID)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}

	return c.JSON(http.StatusOK, nil)
}

func (h *Properties) GetEndpoints() []server.Endpoint {
	return []server.Endpoint{
		server.NewEndpoint(h.GetAll, http.MethodGet, "/users/:uid/properties"),
		server.NewEndpoint(h.GetByID, http.MethodGet, "/users/:uid/properties/:id"),
		server.NewEndpoint(h.Create, http.MethodPost, "/users/properties"),
		server.NewEndpoint(h.DeleteByID, http.MethodDelete, "/users/:uid/properties/:id"),
	}
}
