package Tasks_if

// Doing nothing is very hard to do… you never know when you’re finished.
// Having some tasks can help you a lot by giving you that comfortable feeling that you
// did something.
//

import (
	"context"
	"github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/types"
)

type Repository interface {
	//// List tasks with pagination.
	ListTasks(ctx context.Context, options types.ListingOptions) ([]*taskpb.Task, types.ListingMetas, error)
	//// Returns a single task.
	GetTask(ctx context.Context, idTsk string) (*taskpb.Task, error)
	//// Use this to create new tasks.
	CreateTask(ctx context.Context, data *taskpb.Task) (id string, err error)
	//// Use this to update existing tasks. PATCH is also supported
	UpdateTask(ctx context.Context, data *taskpb.Task, idTsk string) (*taskpb.Task, error)
	//// Use this to delete existing tasks.
	DeleteTask(ctx context.Context, idTsk string) error
	//// Use this to delete ALL tasks.
	DeleteAllTaskss(ctx context.Context) error
	//// Custom methods are always POST.
	SuspendTask(ctx context.Context, idTsk string) error
}
