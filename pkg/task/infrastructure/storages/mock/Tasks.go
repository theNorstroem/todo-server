package mock

// Mock storage CRUD implementation for Tasks

import (
	"context"
	"github.com/theNorstroem/todo-server/pkg/task/errors"
	"github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/searchengine"
	"github.com/veith/fgs-lib/pkg/types"
	"github.com/veith/fgs-lib/pkg/ulid"
)

// CreateTask creates and indexes a taskpb.Task
// Use this to create new tasks.
func (s *Storage) CreateTask(ctx context.Context, data *taskpb.Task) (id string, err error) {
	newItem := data
	newItem.Id = ulid.GenerateStringULID()

	if err := s.db.Write(CollectionTasks, newItem.Id, newItem); err != nil {
		return "", err
	}
	searchengine.Index(CollectionTasks, newItem.Id, newItem)
	return newItem.Id, nil
}

// DeleteTask deletes a
// Use this to delete existing tasks.
func (s *Storage) DeleteTask(ctx context.Context, idTsk string) error {
	err := s.db.Delete(CollectionTasks, idTsk)
	if err != nil {
		return errors.ErrNotFound
	}
	searchengine.DeleteFromIndex(CollectionTasks, idTsk)
	return nil
}

// DeleteAllTaskss deletes all
// Use this to delete ALL tasks.
func (s *Storage) DeleteAllTaskss(ctx context.Context) error {
	panic("implement me")
}

// Returns a single task.
func (s *Storage) GetTask(ctx context.Context, idTsk string) (*taskpb.Task, error) {

	var item *taskpb.Task

	if err := s.db.Read(CollectionTasks, idTsk, &item); err != nil {
		return item, errors.ErrNotFound
	}

	return item, nil
}

// ListTasks returns paginated list of taskpb.Task
// List tasks with pagination.
func (s *Storage) ListTasks(ctx context.Context, options types.ListingOptions) ([]*taskpb.Task, types.ListingMetas, error) {

	list := []*taskpb.Task{}

	// todo implement options.Fields

	res, err := searchengine.List(CollectionTasks, options, map[string]string{})
	if err == nil {
		for _, hit := range res.Hits {
			item, e := s.GetTask(ctx, hit.ID)
			if e == nil {
				// append to list
				list = append(list, item)
			}
		}
	}

	return list, types.ListingMetas{NumOfRecordsForRequest: uint32(res.Total)}, err
}

// SuspendTask deletes all
// Custom methods are always POST.
func (s *Storage) SuspendTask(ctx context.Context, idTsk string) error {
	panic("implement me")
}

//  UpdateTask updates and reindex a
// Use this to update existing tasks. PATCH is also supported
func (s *Storage) UpdateTask(ctx context.Context, data *taskpb.Task, idTsk string) (*taskpb.Task, error) {

	data.Id = idTsk
	if err := s.db.Write(CollectionTasks, idTsk, data); err != nil {
		return &taskpb.Task{}, err
	}
	searchengine.Index(CollectionTasks, idTsk, data)

	return data, nil
}
