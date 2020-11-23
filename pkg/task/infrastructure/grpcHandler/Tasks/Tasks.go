// Code generated by furoc-gen-server. DO NOT EDIT.
// Hint: Put your implementations in an other file in the same package like this file.
// so you ar able to regenerate the stuff, without overwriting yours.
// versions:

package Tasks_grpc

// Doing nothing is very hard to do… you never know when you’re finished.
// Having some tasks can help you a lot by giving you that comfortable feeling that you
// did something.
//

import (
	"github.com/theNorstroem/todo-server/pkg/task"
	"github.com/theNorstroem/todo-server/pkg/task/core"
	"github.com/theNorstroem/todo-server/pkg/task/interfaces/Tasks"
	"github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/hateoas"
	"google.golang.org/grpc"
)

var HTS = hateoas.NewHTSBuilder("Tasks")

type ServiceServer struct {
	taskpb.UnsafeTasksServer
	Tasks Tasks_if.Service
}

// Register your domainServices here
func GetServiceServer(domainServices core.DomainServices) taskpb.TasksServer {
	var service ServiceServer

	service.Tasks = domainServices.Tasks
	return service
}

func Handler(s *grpc.Server) {
	taskpb.RegisterTasksServer(s, GetServiceServer(task.DomainServices))
}
