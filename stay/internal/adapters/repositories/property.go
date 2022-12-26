package repositories

import (
	"gorm.io/gorm"
	"stay/core/entities"
	"stay/internal/ports"
)

// Property represents the space rent on platforms.
type Property struct {
	*entities.Base
	ID          string `gorm:"primaryKey"`
	UserUID     string `gorm:"not null"`
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Address     string `gorm:"not null"`
	Typ         string `gorm:"not null"`
	Rooms       int    `gorm:"not null"`
}

func NewProperty(p *ports.Property) *Property {
	return &Property{
		UserUID:     p.UserUID,
		Name:        p.Name,
		Description: p.Description,
		Address:     p.Address,
		Typ:         p.Typ,
		Rooms:       p.Rooms,
	}
}

// BeforeCreate will set the ID with uuid generated randomly.
func (p *Property) BeforeCreate(_ *gorm.DB) error {
	p.ID = p.GenerateID()
	return nil
}

func (p *Property) ToDomain() *ports.Property {
	if p == nil {
		return nil
	}

	return &ports.Property{
		ID:          p.ID,
		UserUID:     p.UserUID,
		Name:        p.Name,
		Description: p.Description,
		Address:     p.Address,
		Typ:         p.Typ,
		Rooms:       p.Rooms,
	}
}

func NewListPropertyDomain(properties []*Property) []*ports.Property {
	var list []*ports.Property

	for _, p := range properties {
		list = append(list, p.ToDomain())
	}

	return list
}
