package repositories

import (
	"gorm.io/gorm"
	"stay/core/datasources/sql"
	"stay/internal/ports"
)

type propertyReader struct {
	db *gorm.DB
}

func NewPropertyReader(c sql.Connection) ports.PropertyReader {
	db := c.GetDatabase()
	return &propertyReader{db: db}
}

func (r *propertyReader) GetAll(uid string) ([]*ports.Property, error) {
	var properties []*Property
	err := r.db.
		Where("user_uid = ?", uid).
		Find(&properties).Error

	if err != nil {
		return nil, err
	}

	return NewListPropertyDomain(properties), nil
}

func (r *propertyReader) GetProperty(uid, id string) (*ports.Property, error) {
	var p *Property
	err := r.db.
		Where("user_uid = ?", uid).
		Where("id = ?", id).First(&p).Error

	if err != nil {
		return nil, err
	}

	return p.ToDomain(), nil
}
