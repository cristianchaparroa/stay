package repositories

import (
	"gorm.io/gorm"
	"stay/core/datasources/sql"
	"stay/internal/ports"
)

type userReader struct {
	db *gorm.DB
}

func NewUserReader(c sql.Connection) ports.UserReader {
	return &userReader{db: c.GetDatabase()}
}

func (r *userReader) GetUser(uid string) (*ports.User, error) {
	var user *User

	err := r.db.Where("uid = ?", uid).First(&user).Error
	if err != nil {
		return nil, err
	}

	return user.ToDomain(), nil
}
