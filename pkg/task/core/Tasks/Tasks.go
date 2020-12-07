package Tasks

// Doing nothing is very hard to do… you never know when you’re finished.
// Having some tasks can help you a lot by giving you that comfortable feeling that you
// did something.
//

import (
	"context"
	Tasks_if "github.com/theNorstroem/todo-server/pkg/task/interfaces/Tasks"
	"github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/microbus"
	"github.com/veith/fgs-lib/pkg/types"
)

type service struct {
	repository Tasks_if.Repository
	moduleBus  *microbus.EventBus
}

// Used in application.go to connect with storage adapter
func NewService(r interface{}, bus *microbus.EventBus) Tasks_if.Service {
	return &service{repository: r.(Tasks_if.Repository), moduleBus: bus}
}

//// List tasks with pagination.
func (s *service) ListTasks(ctx context.Context, options types.ListingOptions) ([]*taskpb.Task, types.ListingMetas, error) {
	return s.repository.ListTasks(ctx, options)
}

//// Returns a single task.
func (s *service) GetTask(ctx context.Context, idTsk string) (*taskpb.Task, error) {
	return s.repository.GetTask(ctx, idTsk)
}

//// Use this to create new tasks.
func (s *service) CreateTask(ctx context.Context, data *taskpb.Task) (id string, err error) {
	return s.repository.CreateTask(ctx, data)
}

//// Use this to update existing tasks. PATCH is also supported
func (s *service) UpdateTask(ctx context.Context, data *taskpb.Task, idTsk string) (*taskpb.Task, error) {
	return s.repository.UpdateTask(ctx, data, idTsk)
}

//// Use this to delete existing tasks.
func (s *service) DeleteTask(ctx context.Context, idTsk string) error {
	return s.repository.DeleteTask(ctx, idTsk)
}

//// Use this to delete ALL tasks.
func (s *service) DeleteAllTaskss(ctx context.Context) error {
	return s.repository.DeleteAllTaskss(ctx)
}

//// Custom methods are always POST.
func (s *service) SuspendTask(ctx context.Context, idTsk string) error {
	return s.repository.SuspendTask(ctx, idTsk)
}
