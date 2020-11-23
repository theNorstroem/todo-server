package person_grpcHandler

// Register the grpc handlers for Module person

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	Persons_grpc "github.com/theNorstroem/todo-server/pkg/person/infrastructure/grpcHandler/Persons"
)

func RegisterGrpcApis() {
	grpcServer := infrastructure.Infrastructure.Grpc

	// register the grpc handlers for Persons
	Persons_grpc.Handler(grpcServer)
}
