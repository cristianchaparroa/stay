package handlers

import (
	"fmt"
)

const (
	propertyIDParam  PropertyParam = "id"
	propertyUIDParam PropertyParam = "uid"
)

type PropertyParam string

func (p PropertyParam) ToString() string {
	return fmt.Sprintf("%s", p)
}
