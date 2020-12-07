package Persons_if

// People / persons are important to us, they should do the work for us.
//

import (
	"context"
	"github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/types"
)

type Service interface {
	//// List persons with pagination.
	ListPersons(ctx context.Context, options types.ListingOptions) ([]*personpb.Person, types.ListingMetas, error)
	//// Returns a single person.
	GetPerson(ctx context.Context, idPrs string) (*personpb.Person, error)
	//// Use this to create new persons.
	CreatePerson(ctx context.Context, data *personpb.Person) (id string, err error)
	//// Use this to update existing persons. PATCH is also supported
	UpdatePerson(ctx context.Context, data *personpb.Person, idPrs string) (*personpb.Person, error)
	//// We use this to disable a person in the list, we do not delete them.
	DeletePerson(ctx context.Context, idPrs string) error
	//// Fireing some persons can increase the performance of the other persons POST. Do not use this to much. The big downside is, that you have to assign their work to somone other.
	FirePerson(ctx context.Context, idPrs string) error
}
