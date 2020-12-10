package Tasks_grpc

// patches an item with values from the request body and the update_mask

import (
	"github.com/theNorstroem/todo-specs/dist/pb/task"
)

func PatchWithUpdateMask(item *taskpb.Task, request *taskpb.UpdateTaskRequest) *taskpb.Task {
	// bypass if no update_mask was given (this must be a PUT method on the gateway)
	if request.UpdateMask != nil {
		// from spec
		in := request.Body
		for _, field := range request.UpdateMask.Paths {

			if "id" == field {
				item.Id = in.Id
			}

			if "display_name" == field {
				item.DisplayName = in.DisplayName
			}

			if "note.value" == field {
				item.Note.Value = in.Note.Value
			}

			if "parent.id" == field {
				item.Parent.Id = in.Parent.Id
			}

			if "parent.display_name" == field {
				item.Parent.DisplayName = in.Parent.DisplayName
			}

			if "checklist" == field {
				item.Checklist = in.Checklist
			}

			if "due_date.year" == field {
				item.DueDate.Year = in.DueDate.Year
			}

			if "due_date.month" == field {
				item.DueDate.Month = in.DueDate.Month
			}

			if "due_date.day" == field {
				item.DueDate.Day = in.DueDate.Day
			}

			if "related_tasks" == field {
				item.RelatedTasks = in.RelatedTasks
			}

			if "involved_persons" == field {
				item.InvolvedPersons = in.InvolvedPersons
			}

			if "responsible_person.id" == field {
				item.ResponsiblePerson = in.ResponsiblePerson
			}

			if "responsible_person.display_name" == field {
				item.ResponsiblePerson.DisplayName = in.ResponsiblePerson.DisplayName
			}

			if "done" == field {
				item.Done = in.Done
			}

		}
	} else {
		item = request.Body
	}

	return item
}
