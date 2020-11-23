package upper

import "github.com/oklog/ulid"

type User struct {
	Ulid     ulid.ULID `json:"ulid,omitempty" db:"ulid,omitempty"`
	Username string    `json:"username,omitempty" db:"username,omitempty"`
	Password string    `json:"password,omitempty" db:"password,omitempty"`
	Rolle    string    `json:"rolle,omitempty" db:"rolle,omitempty"`
}
