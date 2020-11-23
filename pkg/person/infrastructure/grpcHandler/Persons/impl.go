package Persons_grpc

// Minimal CRUD implementation for Persons

import (
	"context"
	"github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/signatures"
	"github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/hateoas"
	"github.com/veith/fgs-lib/pkg/pagination"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreatePerson handles the request that comes on:
// /persons
// Use this to create new persons.
func (s ServiceServer) CreatePerson(ctx context.Context, request *personpb.CreatePersonRequest) (*signaturespb.EmptyEntity, error) {
	id, err := s.Persons.CreatePerson(ctx, request.Body)
	if err != nil {
		return nil, err
	}
	var entity = &signaturespb.EmptyEntity{}

	entity.Links = HTS.EntityHTS(
		ctx,
		"person.Person",
		hateoas.Idmap{},
		"/persons/"+id,
		"DELETE", "PATCH",
	)
	return entity, err
}

// DeletePerson handles the delete request that comes on:
// /persons/{pn}
// We use this to disable a person in the list, we do not delete them.
func (s ServiceServer) DeletePerson(ctx context.Context, request *personpb.DeletePersonRequest) (*emptypb.Empty, error) {

	err := s.Persons.DeletePerson(ctx, request.Prs)

	_ = grpc.SendHeader(ctx, metadata.Pairs(
		"grpc-status", "204 No Content",
	))

	return &emptypb.Empty{}, err
}

// FirePerson handles the delete request that comes on:
// URL: /persons/{pn}:ferment
// Fireing some persons can increase the performance of the other persons POST. Do not use this to much.
func (s ServiceServer) FirePerson(ctx context.Context, request *personpb.FirePersonRequest) (*emptypb.Empty, error) {
	// no autogeneration at the moment :-(
	panic("implement me")
}

// GetPerson handles the get request that comes on:
// /persons/{pn}
// Returns a single person.
func (s ServiceServer) GetPerson(ctx context.Context, request *personpb.GetPersonRequest) (*personpb.PersonEntity, error) {

	var entity *personpb.PersonEntity

	item, err := s.Persons.GetPerson(ctx, request.Prs)
	if err != nil {
		return entity, err
	}

	entity = &personpb.PersonEntity{
		Data: item,
		Links: HTS.EntityHTS(
			ctx,
			"person.Person",
			hateoas.Idmap{"pn": item.Id},
			"/persons/{pn}",
			"DELETE", "PATCH",
		),
	}

	return entity, err
}

// ListPersons handles the list request that comes on:
// /persons
// List persons with pagination.
func (s ServiceServer) ListPersons(ctx context.Context, request *personpb.ListPersonsRequest) (*personpb.PersonCollection, error) {

	entities := []*personpb.PersonEntity{}
	list, listingMetas, _ := s.Persons.ListPersons(ctx, pagination.GetListingOptions(request))

	for _, item := range list {
		entry := item
		entities = append(entities,
			&personpb.PersonEntity{
				Data: entry,
				Links: HTS.EntityHTS(
					ctx,
					"person.Person",
					hateoas.Idmap{"pn": item.Id},
					"/persons/{pn}",
					"DELETE", "PATCH",
				),
			})
	}

	collection := &personpb.PersonCollection{Entities: entities}
	collection.Links = HTS.CollectionHTS(ctx,
		pagination.GetListingOptions(request),
		listingMetas,
		hateoas.Idmap{},
		"/persons",
		"CREATE")

	return collection, nil
}

// UpdatePerson handles the update request that comes on:
// PUT /persons/{pn}
// Use this to update existing persons. PATCH is also supported
func (s ServiceServer) UpdatePerson(ctx context.Context, request *personpb.UpdatePersonRequest) (*personpb.PersonEntity, error) {

	// read the original, because we can patch
	item, err := s.Persons.GetPerson(ctx, request.Prs)
	if err != nil {
		return nil, err
	}

	// patch fields according to update_mask
	item = PatchWithUpdateMask(item, request)

	// save the changes
	item, err = s.Persons.UpdatePerson(ctx, item, request.Prs)

	if err != nil {
		return nil, err
	}

	return &personpb.PersonEntity{
		Data: item,
		Links: HTS.EntityHTS(
			ctx,
			"person.Person",
			hateoas.Idmap{"pn": request.Prs},
			"/persons/{pn}",
			"DELETE", "PATCH",
		),
	}, err
}
