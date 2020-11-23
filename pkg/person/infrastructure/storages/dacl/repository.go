package dacl

// Data Access Layer storage for person

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	Persons_if "github.com/theNorstroem/todo-server/pkg/person/interfaces/Persons"
	"github.com/upper/db/v4"
)

const (
	// CollectionName for the dacl storage
	CollectionPersons = "person.Persons"
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

func (s *Storage) Persons() Persons_if.Repository {
	// Hint: implement s in a separate file
	// so a re generation would not overwrite your implementation
	return s
}
