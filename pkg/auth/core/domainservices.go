package core

import (
	"github.com/theNorstroem/todo-server/pkg/auth/core/login"
	//"github.com/theNorstroem/todo-server/pkg/auth/core/logout"
)

// Add your DomainServices to this struct.
type DomainServices struct {
	Login login.Service
	//Logout logout.Service
}
