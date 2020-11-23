package upper

import (
	"github.com/theNorstroem/todo-server/pkg/auth/core/login"
	"github.com/upper/db/v4"
	"github.com/veith/fgs-lib/pkg/auth"
)

func (s *Storage) GetProfile(userId string) (credentials *auth.Profile, err error) {
	var profile *auth.Profile
	var user User

	res := s.dbCollection.Find(db.Cond{"username": userId})
	if err := res.One(&user); err != nil {
		return profile, login.ErrNotFound(err.Error())
	}

	profile = user.ToProfile()

	return profile, nil
}

func (u *User) ToProfile() (credentials *auth.Profile) {

	if credentials == nil {
		credentials = new(auth.Profile)
	}
	credentials.Username = u.Username
	credentials.Pwhash = u.Password

	if credentials.Roles == nil {
		credentials.Roles = []string{}
	}
	credentials.Roles = append(credentials.Roles, u.Rolle)

	return credentials
}
