package Tasks_grpc

import (
	"context"
	taskpb "github.com/theNorstroem/todo-specs/dist/pb/task"
	"github.com/veith/fgs-lib/pkg/errors"
)

func validateAgainstSpec(ctx context.Context, data *taskpb.Task) error {
	e := errors.ErrConstraintViolation()
	// required
	if data.DisplayName == "" {
		e.AddFieldViolation("display_name", "A display_name is required.")
	}
	/**
		// min is 2
		if len(data.InvolvedPersons) < 2 {
			e.AddFieldViolation("involved_persons.0", "Somebody else must be involved.")
		}
		// max is 4
		if len(data.InvolvedPersons) > 4 {
			e.AddFieldViolation("involved_persons", "Too many people.")
		}
	**/
	// return error if we had any violation
	if e.HasFieldViolation() {
		return e.GetErr()
	}
	return nil
}
