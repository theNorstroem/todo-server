package Persons

// People / persons are important to us, they should do the work for us.
//

import (
	"context"
	Persons_if "github.com/theNorstroem/todo-server/pkg/person/interfaces/Persons"
	"github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/microbus"
	"github.com/veith/fgs-lib/pkg/types"
)

type service struct {
	repository Persons_if.Repository
	moduleBus  *microbus.EventBus
}

// Used in application.go to connect with storage adapter
func NewService(r interface{}, bus *microbus.EventBus) Persons_if.Service {
	return &service{repository: r.(Persons_if.Repository), moduleBus: bus}
}

//// List persons with pagination.
func (s *service) ListPersons(ctx context.Context, options types.ListingOptions) ([]*personpb.Person, types.ListingMetas, error) {
	return s.repository.ListPersons(ctx, options)
}

//// Returns a single person.
func (s *service) GetPerson(ctx context.Context, idPrs string) (*personpb.Person, error) {
	return s.repository.GetPerson(ctx, idPrs)
}

//// Use this to create new persons.
func (s *service) CreatePerson(ctx context.Context, data *personpb.Person) (id string, err error) {
	return s.repository.CreatePerson(ctx, data)
}

//// Use this to update existing persons. PATCH is also supported
func (s *service) UpdatePerson(ctx context.Context, data *personpb.Person, idPrs string) (*personpb.Person, error) {
	return s.repository.UpdatePerson(ctx, data, idPrs)
}

//// We use this to disable a person in the list, we do not delete them.
func (s *service) DeletePerson(ctx context.Context, idPrs string) error {
	return s.repository.DeletePerson(ctx, idPrs)
}

//// Fireing some persons can increase the performance of the other persons POST. Do not use this to much. The big downside is, that you have to assign their work to somone other.
func (s *service) FirePerson(ctx context.Context, idPrs string) error {
	return s.repository.FirePerson(ctx, idPrs)
}
