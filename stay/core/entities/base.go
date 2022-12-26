package entities

import uuid "github.com/satori/go.uuid"

type Base struct{}

func (b *Base) GenerateID() string {
	uuid := uuid.NewV4()
	return uuid.String()
}
