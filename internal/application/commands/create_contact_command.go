package commands

import (
	"example.com/go_agenda/internal/application/value_objects"
	"example.com/go_agenda/internal/domain/entities"
)

type CreateContactCommand struct{}

func (createContactCommand CreateContactCommand) Execute(command interface{}) (interface{}, error) {
	// create external id in the core layer
	contactToCreate := command.(value_objects.ContactToCreate)
	core_contact := entities.Contact{
		Name:  contactToCreate.Name,
		Phone: contactToCreate.Phone,
		Email: contactToCreate.Email,
	}
	core_contact.GenerateExternalId()

	// save on database

	// build result
	contact_result := value_objects.CreatedContact{
		Id:    core_contact.ExternalId.String(),
		Name:  core_contact.Name,
		Phone: core_contact.Phone,
		Email: core_contact.Email,
	}
	return contact_result, nil
}
