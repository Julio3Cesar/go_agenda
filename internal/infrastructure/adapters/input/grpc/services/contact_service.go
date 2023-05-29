package grpc_services

import (
	"context"
	"log"

	ports_input "example.com/go_agenda/internal/application/ports/input"
	"example.com/go_agenda/internal/application/value_objects"
	pb "example.com/go_agenda/internal/infrastructure/adapters/input/grpc/protocol_buffers"
)

type ContactService struct {
	pb.UnimplementedContactServiceServer
	CreateContactCommand ports_input.Command
}

func (contactService ContactService) Create(ctx context.Context, contact *pb.Contact) (*pb.CreatedContact, error) {
	log.Printf("Trying to create the contact: %v", contact)

	contactToCreate := value_objects.ContactToCreate{
		Name:  contact.Name,
		Phone: contact.Phone,
		Email: contact.Email,
	}
	created_contact, error := contactService.CreateContactCommand.Execute(contactToCreate)
	if error != nil {
		log.Fatalf("Error on create contact: %v", error)
	}

	created_contact_result := created_contact.(value_objects.CreatedContact)
	return &pb.CreatedContact{Id: created_contact_result.Id, Name: created_contact_result.Name, Phone: created_contact_result.Phone, Email: created_contact_result.Email}, nil
}
