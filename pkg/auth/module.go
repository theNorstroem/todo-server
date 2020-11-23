package auth

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	"github.com/theNorstroem/todo-server/pkg/auth/apis/grpc"
	"github.com/theNorstroem/todo-server/pkg/auth/core"
	"github.com/theNorstroem/todo-server/pkg/auth/core/login"
	"github.com/theNorstroem/todo-server/pkg/auth/storage/json"
	"github.com/theNorstroem/todo-server/pkg/auth/storage/upper"
)

// StorageType defines available storage types
type Type int

const (
	// JSON will store data in JSON files saved on disk
	JSON Type = iota
	// Memory will store data in memory
	Memory
	// Upper-DB will store data in a Database
	Upper
)

func RegisterModule() {
	// restRouter := infrastructure.AdminRest
	grpcServer := infrastructure.Infrastructure.Grpc

	// set up storage
	storageType := JSON // this could be a flag; hardcoded here for simplicity
	var authServices core.DomainServices

	switch storageType {
	case JSON:
		// error handling omitted for simplicity
		s, _ := json.NewStorage()
		authServices.Login = login.NewService(s)
		//authServices.Logout = logout.NewService(s)
	case Upper:
		s := upper.NewStorage()
		authServices.Login = login.NewService(s)
	}

	// register the handlers
	//-----------------------

	// register the grpc handlers
	grpc.Handler(grpcServer, authServices)

}
