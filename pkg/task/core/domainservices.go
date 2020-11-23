package core

// Domainservices for Module task

import (
	Tasks_if "github.com/theNorstroem/todo-server/pkg/task/interfaces/Tasks"
)

// Add your DomainServices to this struct.
type DomainServices struct {
	Tasks Tasks_if.Service
}
