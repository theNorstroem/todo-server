package core

// Domainservices for Module person

import (
	Persons_if "github.com/theNorstroem/todo-server/pkg/person/interfaces/Persons"
)

// Add your DomainServices to this struct.
type DomainServices struct {
	Persons Persons_if.Service
}
