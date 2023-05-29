package grpc_servers

import (
	"log"
	"net"

	pb "example.com/go_agenda/internal/infrastructure/adapters/input/grpc/protocol_buffers"
	"google.golang.org/grpc"
)

type AgendaGrpcServer struct {
	ContactService pb.ContactServiceServer
}

const port = ":9000"

func (agendaGrpcServer AgendaGrpcServer) Run() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on port %v: %v", port, err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterContactServiceServer(grpcServer, agendaGrpcServer.ContactService)
	log.Printf("Server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
