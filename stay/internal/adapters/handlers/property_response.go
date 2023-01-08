package handlers

import "stay/internal/ports"

type PropertyResponse struct {
	ID          string `json:"id"`
	UserUID     string `json:"user_uid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Address     string `json:"address"`
	Typ         string `json:"typ"`
	Rooms       int    `json:"rooms"`
}

func NewPropertyResponse(p *ports.Property) *PropertyResponse {
	return &PropertyResponse{
		ID:          p.ID,
		UserUID:     p.UserUID,
		Name:        p.Name,
		Description: p.Description,
		Address:     p.Address,
		Typ:         p.Typ,
		Rooms:       p.Rooms,
	}
}

func NewPropertyListResponse(list []*ports.Property) []*PropertyResponse {
	response := make([]*PropertyResponse, 0)

	for _, p := range list {
		response = append(response, NewPropertyResponse(p))
	}

	return response
}
