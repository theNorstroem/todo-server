package task_grpcHandler

// Register the grpc handlers for Module task

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	Tasks_grpc "github.com/theNorstroem/todo-server/pkg/task/infrastructure/grpcHandler/Tasks"
)

func RegisterGrpcApis() {
	grpcServer := infrastructure.Infrastructure.Grpc

	// register the grpc handlers for Tasks
	Tasks_grpc.Handler(grpcServer)
}
