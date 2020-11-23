package login

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// The user could not be found
func ErrNotFound(Errormessage string) error {
	return status.New(codes.NotFound, Errormessage).Err()
}

func LoginFailed(Errormessage string) error {
	return status.New(codes.Unauthenticated, Errormessage).Err()
}
