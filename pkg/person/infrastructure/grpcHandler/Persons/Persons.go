// Code generated by furoc-gen-server. DO NOT EDIT.
// Hint: Put your implementations in an other file in the same package like this file.
// so you ar able to regenerate the stuff, without overwriting yours.
// versions:

package Persons_grpc

// People / persons are important to us, they should do the work for us.
//

import (
	"github.com/theNorstroem/todo-server/pkg/person"
	"github.com/theNorstroem/todo-server/pkg/person/core"
	"github.com/theNorstroem/todo-server/pkg/person/interfaces/Persons"
	"github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/hateoas"
	"google.golang.org/grpc"
)

var HTS = hateoas.NewHTSBuilder("Persons")

type ServiceServer struct {
	personpb.UnsafePersonsServer
	Persons Persons_if.Service
}

// Register your domainServices here
func GetServiceServer(domainServices core.DomainServices) personpb.PersonsServer {
	var service ServiceServer

	service.Persons = domainServices.Persons
	return service
}

func Handler(s *grpc.Server) {
	personpb.RegisterPersonsServer(s, GetServiceServer(person.DomainServices))
}
