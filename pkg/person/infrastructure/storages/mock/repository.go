package mock

// Mock storage for person

import (
	"github.com/sdomino/scribble"
	"github.com/spf13/viper"
	Persons_if "github.com/theNorstroem/todo-server/pkg/person/interfaces/Persons"
	"os"
)

const (
	// CollectionName for the mock storage
	// use it in s.db.Write( CollectionName, data.Id, data);
	CollectionPersons = "person.Persons"
)

// Storage stores  data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new File  storage
func NewStorage() (*Storage, error) {
	var err error
	s := new(Storage)
	cwd, _ := os.Getwd()
	s.db, err = scribble.New(cwd+viper.GetString("storage.basedir"), nil)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) Persons() Persons_if.Repository {
	// Hint: implement s in a separate file
	// so a re generation would not overwrite your implementation
	return s
}
