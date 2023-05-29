package main

import (
	"context"
	"log"
	"time"

	pb "example.com/go_agenda/internal/infrastructure/adapters/input/grpc/protocol_buffers"
	"google.golang.org/grpc"
)

const address = "localhost:9000"

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	contactClient := pb.NewContactServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	createdContact, err := contactClient.Create(ctx, &pb.Contact{Name: "Julio", Phone: "9999", Email: "mail@mail.com"})
	if err != nil {
		log.Fatalf("Could not create the contact: %v", err)
	}

	log.Printf("CreatedContact: %v", createdContact)
}
