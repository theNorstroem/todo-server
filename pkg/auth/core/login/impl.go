package login

import (
	authpb "github.com/theNorstroem/todo-specs/dist/pb/auth"
	"github.com/veith/fgs-lib/pkg/auth"
)

// send the auth request to the repository
func (s service) Login(credentials authpb.Credentials) (token string, err error) {
	var profile *auth.Profile
	profile, err = s.repository.GetProfile(credentials.Username) // error handling omitted for simplicity
	if err != nil {
		return token, err
	}

	if auth.CheckPasswordHash(credentials.Password, profile.Pwhash) {
		// any validation can be done here
		// jwt stuff and so on

		token = auth.CreateJWT(profile)
		return token, nil

	} else {
		err = LoginFailed("incorrect password or username")
	}

	return token, err
}
