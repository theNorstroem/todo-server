package mock

// Mock storage CRUD implementation for Persons

import (
	"context"
	"github.com/theNorstroem/todo-server/pkg/person/errors"
	"github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/searchengine"
	"github.com/veith/fgs-lib/pkg/types"
	"github.com/veith/fgs-lib/pkg/ulid"
)

// CreatePerson creates and indexes a personpb.Person
// Use this to create new persons.
func (s *Storage) CreatePerson(ctx context.Context, data *personpb.Person) (id string, err error) {
	newItem := data
	newItem.Id = ulid.GenerateStringULID()

	if err := s.db.Write(CollectionPersons, newItem.Id, newItem); err != nil {
		return "", err
	}
	searchengine.Index(CollectionPersons, newItem.Id, newItem)
	return newItem.Id, nil
}

// DeletePerson deletes a
// We use this to disable a person in the list, we do not delete them.
func (s *Storage) DeletePerson(ctx context.Context, idPrs string) error {
	err := s.db.Delete(CollectionPersons, idPrs)
	if err != nil {
		return errors.ErrNotFound
	}
	searchengine.DeleteFromIndex(CollectionPersons, idPrs)
	return nil
}

// FirePerson deletes all
// Fireing some persons can increase the performance of the other persons POST. Do not use this to much. The big downside is, that you have to assign their work to somone other.
func (s *Storage) FirePerson(ctx context.Context, idPrs string) error {
	panic("implement me")
}

// Returns a single person.
func (s *Storage) GetPerson(ctx context.Context, idPrs string) (*personpb.Person, error) {

	var item *personpb.Person

	if err := s.db.Read(CollectionPersons, idPrs, &item); err != nil {
		return item, errors.ErrNotFound
	}

	return item, nil
}

// ListPersons returns paginated list of personpb.Person
// List persons with pagination.
func (s *Storage) ListPersons(ctx context.Context, options types.ListingOptions) ([]*personpb.Person, types.ListingMetas, error) {

	list := []*personpb.Person{}

	// todo implement options.Fields

	res, err := searchengine.List(CollectionPersons, options, map[string]string{})
	if err == nil {
		for _, hit := range res.Hits {
			item, e := s.GetPerson(ctx, hit.ID)
			if e == nil {
				// append to list
				list = append(list, item)
			}
		}
	}

	return list, types.ListingMetas{NumOfRecordsForRequest: uint32(res.Total)}, err
}

//  UpdatePerson updates and reindex a
// Use this to update existing persons. PATCH is also supported
func (s *Storage) UpdatePerson(ctx context.Context, data *personpb.Person, idPrs string) (*personpb.Person, error) {

	data.Id = idPrs
	if err := s.db.Write(CollectionPersons, idPrs, data); err != nil {
		return &personpb.Person{}, err
	}
	searchengine.Index(CollectionPersons, idPrs, data)

	return data, nil
}
