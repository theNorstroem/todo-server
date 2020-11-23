package login

import (
	authpb "github.com/theNorstroem/todo-specs/dist/pb/auth"
	"github.com/veith/fgs-lib/pkg/auth"
)

// Service provides auth adding operations.
type Service interface {
	// send the credentials to log in
	Login(authpb.Credentials) (token string, err error)
}

// Repository provides access to auth repository.
type Repository interface {
	// Login, userid could be the username or so
	GetProfile(userId string) (profile *auth.Profile, err error)
}

type service struct {
	repository Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}
