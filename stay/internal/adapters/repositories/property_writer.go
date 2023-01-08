package repositories

import (
	"gorm.io/gorm"
	"stay/core/datasources/sql"
	"stay/internal/ports"
)

type propertyWriter struct {
	db *gorm.DB
}

func NewPropertyWriter(c sql.Connection) ports.PropertyWriter {
	return &propertyWriter{db: c.GetDatabase()}
}

func (w *propertyWriter) Save(p *ports.Property) (*ports.Property, error) {
	property := NewProperty(p)
	tx := w.db.Save(property)

	if tx.Error != nil {
		return nil, tx.Error
	}
	return property.ToDomain(), nil
}

func (w *propertyWriter) Delete(id string) error {
	tx := w.db.Where("id = ?", id).Delete(&Property{})

	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
