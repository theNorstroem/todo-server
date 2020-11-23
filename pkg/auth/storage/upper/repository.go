package upper

import (
	"github.com/theNorstroem/todo-server/internal/infrastructure"
	"github.com/upper/db/v4"
)

const (
	// Collection identifier for the Upper collection of users
	Collection = "sys_user"
)

// Storage stores sample data in JSON files
type Storage struct {
	db           db.Session
	dbCollection db.Collection
}

// NewStorage returns a new DB  storage
func NewStorage() *Storage {
	s := new(Storage)
	s.db = infrastructure.Infrastructure.DbSession
	s.dbCollection = s.db.Collection(Collection)
	return s
}
