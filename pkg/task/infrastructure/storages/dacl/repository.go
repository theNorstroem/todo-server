package dacl

// Data Access Layer storage for task

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	Tasks_if "github.com/theNorstroem/todo-server/pkg/task/interfaces/Tasks"
	"github.com/upper/db/v4"
)

const (
	// CollectionName for the dacl storage
	CollectionTasks = "task.Tasks"
)

// Storage stores in Database
type Storage struct {
	db db.Session
}

// NewStorage returns a new DB storage
func NewStorage() *Storage {
	s := new(Storage)
	s.db = infrastructure.Infrastructure.DbSession
	return s
}

func (s *Storage) Tasks() Tasks_if.Repository {
	// Hint: implement s in a separate file
	// so a re generation would not overwrite your implementation
	return s
}
