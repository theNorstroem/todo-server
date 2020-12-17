package Tasks_grpc

// Minimal CRUD implementation for Tasks

import (
	"context"
	"encoding/base64"
	"encoding/json"
	filterpb "github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/filter"
	"github.com/theNorstroem/FuroBaseSpecs/dist/pb/furo/signatures"
	"github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/hateoas"
	"github.com/veith/fgs-lib/pkg/pagination"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
)

// CreateTask handles the request that comes on:
// /tasks
// Use this to create new tasks.
func (s ServiceServer) CreateTask(ctx context.Context, request *taskpb.CreateTaskRequest) (*signaturespb.EmptyEntity, error) {
	id, err := s.Tasks.CreateTask(ctx, request.Body)
	if err != nil {
		return nil, err
	}
	var entity = &signaturespb.EmptyEntity{}

	entity.Links = HTS.EntityHTS(
		ctx,
		"task.Task",
		hateoas.Idmap{},
		"/tasks/"+id,
		"DELETE", "PATCH",
	)
	return entity, err
}

// DeleteTask handles the delete request that comes on:
// /tasks/{tsk}
// Use this to delete existing tasks.
func (s ServiceServer) DeleteTask(ctx context.Context, request *taskpb.DeleteTaskRequest) (*emptypb.Empty, error) {

	err := s.Tasks.DeleteTask(ctx, request.Tsk)

	_ = grpc.SendHeader(ctx, metadata.Pairs(
		"grpc-status", "204 No Content",
	))

	return &emptypb.Empty{}, err
}

// DeleteAllTaskss handles the delete request that comes on:
// URL: /tasks
// Use this to delete ALL tasks.
func (s ServiceServer) DeleteAllTaskss(ctx context.Context, request *taskpb.DeleteAllTaskssRequest) (*emptypb.Empty, error) {
	// no autogeneration at the moment :-(
	panic("implement me")
}

// GetTask handles the get request that comes on:
// /tasks/{tsk}
// Returns a single task.
func (s ServiceServer) GetTask(ctx context.Context, request *taskpb.GetTaskRequest) (*taskpb.TaskEntity, error) {

	var entity *taskpb.TaskEntity

	item, err := s.Tasks.GetTask(ctx, request.Tsk)
	if err != nil {
		return entity, err
	}

	entity = &taskpb.TaskEntity{
		Data: item,
		Links: HTS.EntityHTS(
			ctx,
			"task.Task",
			hateoas.Idmap{"tsk": item.Id},
			"/tasks/{tsk}",
			"DELETE", "PATCH",
		),
	}

	return entity, err
}

// ListTasks handles the list request that comes on:
// /tasks
// List tasks with pagination.
func (s ServiceServer) ListTasks(ctx context.Context, request *taskpb.ListTasksRequest) (*taskpb.TaskCollection, error) {

	entities := []*taskpb.TaskEntity{}
	list, listingMetas, _ := s.Tasks.ListTasks(ctx, pagination.GetListingOptions(request))

	filter := &taskpb.Filter{}
	base, _ := base64.StdEncoding.DecodeString(request.Filter)
	json.Unmarshal(base, filter)
	filteroption := filterpb.Filter{}
	filteroption.Clause = []*filterpb.Condition{}

	if filter.Person != nil {
		filteroption.Clause = append(filteroption.Clause, filter.Person)
	}

	if filter.DueDate != nil {
		filteroption.Clause = append(filteroption.Clause, filter.DueDate)
	}

	if filter.Note != nil {
		filteroption.Clause = append(filteroption.Clause, filter.Note)
	}

	for _, item := range list {
		entry := item
		entities = append(entities,
			&taskpb.TaskEntity{
				Data: entry,
				Links: HTS.EntityHTS(
					ctx,
					"task.Task",
					hateoas.Idmap{"tsk": item.Id},
					"/tasks/{tsk}",
					"DELETE", "PATCH",
				),
			})
	}

	collection := &taskpb.TaskCollection{Entities: entities}
	collection.Links = HTS.CollectionHTS(ctx,
		pagination.GetListingOptions(request),
		listingMetas,
		hateoas.Idmap{},
		"/tasks",
		"CREATE")

	return collection, nil
}

// SuspendTask handles the delete request that comes on:
// URL: /tasks/{tsk}:suspend
// Custom methods are always POST.
func (s ServiceServer) SuspendTask(ctx context.Context, request *taskpb.SuspendTaskRequest) (*emptypb.Empty, error) {
	// no autogeneration at the moment :-(
	panic("implement me")
}

// UpdateTask handles the update request that comes on:
// PUT /tasks/{tsk}
// Use this to update existing tasks. PATCH is also supported
func (s ServiceServer) UpdateTask(ctx context.Context, request *taskpb.UpdateTaskRequest) (*taskpb.TaskEntity, error) {

	// read the original, because we can patch
	item, err := s.Tasks.GetTask(ctx, request.Tsk)
	if err != nil {
		return nil, err
	}

	// patch fields according to update_mask
	item = PatchWithUpdateMask(item, request)

	// save the changes
	item, err = s.Tasks.UpdateTask(ctx, item, request.Tsk)

	if err != nil {
		return nil, err
	}

	return &taskpb.TaskEntity{
		Data: item,
		Links: HTS.EntityHTS(
			ctx,
			"task.Task",
			hateoas.Idmap{"tsk": request.Tsk},
			"/tasks/{tsk}",
			"DELETE", "PATCH",
		),
	}, err
}
