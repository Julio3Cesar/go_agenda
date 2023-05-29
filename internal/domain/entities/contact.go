package entities

import (
	"github.com/google/uuid"
)

type Contact struct {
	id         int32
	Name       string
	Phone      string
	Email      string
	ExternalId uuid.UUID
}

func (contact *Contact) GenerateExternalId() {
	contact.ExternalId = uuid.New()
}
