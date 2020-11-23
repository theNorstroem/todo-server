package task

// This file is the initiator for the module task

import (
	"github.com/theNorstroem/todo-server/pkg/task/core"
	Tasks "github.com/theNorstroem/todo-server/pkg/task/core/Tasks"
	// "github.com/theNorstroem/todo-server/pkg/task/infrastructure/eventhandler"
	"github.com/spf13/viper"
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	// "github.com/theNorstroem/todo-server/pkg/task/infrastructure/storages/dacl"
	"github.com/theNorstroem/todo-server/pkg/task/infrastructure/storages/mock"
)

var DomainServices = core.DomainServices{}

func RegisterApplicationModule() {

	//microbus (bus in core model only)
	//var messagebus = microbus.NewMicrobus()

	// use the system bus
	var messagebus = infrastructure.Infrastructure.InternalEventBus

	var storage interface{}
	// set up storage
	storageType := viper.GetString("modules.task.repository") // configure this in the main config

	switch storageType {
	case "TEST": // TEST will store data in memory (for testing)
		// storage = new(memory.Storage)
		break
	case "DACL": // DACL (data access layer) can store in different databases (with upper)
		// storage = dacl.NewStorage()
		break
	case "MOCK": // MOCK will store data in JSON files saved on disk
		storage, _ = mock.NewStorage(messagebus)
		break
	default:
		storage, _ = mock.NewStorage(messagebus)
	}

	// connect the storage handlers for Tasks
	DomainServices.Tasks = Tasks.NewService(storage, messagebus)

	//eventhandler.RegisterSubscriptions(messagebus, &DomainServices)

}
