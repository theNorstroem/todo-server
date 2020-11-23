package Persons_grpc

// patches an item with values from the request body and the update_mask

import (
	"github.com/theNorstroem/todo-specs/dist/pb/person"
)

func PatchWithUpdateMask(item *personpb.Person, request *personpb.UpdatePersonRequest) *personpb.Person {
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

			if "note" == field {
				item.Note = in.Note
			}

			if "email" == field {
				item.Email = in.Email
			}

			if "mobile" == field {
				item.Mobile = in.Mobile
			}

		}
	} else {
		item = request.Body
	}

	return item
}
