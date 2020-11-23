package json

import (
	"github.com/nanobox-io/golang-scribble"
	"github.com/spf13/viper"
	"github.com/theNorstroem/todo-server/pkg/auth/core/login"
	"github.com/veith/fgs-lib/pkg/auth"
	"os"
)

const (

	// CollectionExcerciseDatasheet identifier for the JSON collection of excerciseDatasheets
	UserCollection = "users"
)

// Storage stores excerciseDatasheet data in JSON files
type Storage struct {
	db *scribble.Driver
}

func (s *Storage) GetProfile(userId string) (credentials *auth.Profile, err error) {
	var profile *auth.Profile

	if err := s.db.Read(UserCollection, userId, &profile); err != nil {
		// err handling omitted for simplicity
		return profile, login.ErrNotFound(err.Error())
	}

	return profile, nil
}

// NewStorage returns a new JSON  storage
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
