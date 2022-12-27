package domain

import (
	"errors"
	"stay/internal/ports"
)

type propertyUseCase struct {
	userReader ports.UserReader
	reader     ports.PropertyReader
	writer     ports.PropertyWriter
}

func NewPropertyUseCase(u ports.UserReader, w ports.PropertyWriter, r ports.PropertyReader) ports.PropertyUseCase {
	return &propertyUseCase{
		userReader: u,
		reader:     r,
		writer:     w,
	}
}

func (c *propertyUseCase) GetAll(uid string) ([]*ports.Property, error) {
	_, err := c.userReader.GetUser(uid)
	if err != nil {
		return nil, errors.New(errUserNotFound)
	}

	return c.reader.GetAll(uid)
}

func (c *propertyUseCase) GetProperty(uid, id string) (*ports.Property, error) {
	_, err := c.userReader.GetUser(uid)
	if err != nil {
		return nil, errors.New(errUserNotFound)
	}

	return c.reader.GetProperty(uid, id)
}

func (c *propertyUseCase) Save(p *ports.Property) (*ports.Property, error) {
	_, err := c.userReader.GetUser(p.UserUID)
	if err != nil {
		return nil, errors.New(errUserNotFound)
	}

	return c.writer.Save(p)
}
