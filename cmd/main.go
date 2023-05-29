package main

import (
	"example.com/go_agenda/internal/application/commands"
	grpc_servers "example.com/go_agenda/internal/infrastructure/adapters/input/grpc"
	grpc_services "example.com/go_agenda/internal/infrastructure/adapters/input/grpc/services"
)

func main() {
	contactService := grpc_services.ContactService{
		CreateContactCommand: commands.CreateContactCommand{},
	}

	grpc_servers.AgendaGrpcServer{ContactService: contactService}.Run()
}
