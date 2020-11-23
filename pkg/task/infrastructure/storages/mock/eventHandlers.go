/**
This file is here to show, that even storages can register to the event bus.
This can make sense for file based storages to force a sync or update some of the files on some conditions.

The mock server uses json files to store data, references which are used are not updated when the referenced object
itself changes.

*/
package mock

import (
	"fmt"
	personpb "github.com/theNorstroem/todo-specs/dist/pb/person"
	"github.com/veith/fgs-lib/pkg/microbus"
	"time"
)

// Add your subcriptions to the bus
func registerSubscriptions(bus *microbus.EventBus) {
	// DO NOT  make multiple funcs per topic possible
	// you can wrap the multiple functions in a function that calls the other ones, if needed
	var subscriptionList = map[string]func(event microbus.DataEvent){

		"person.Persons.Updated": updatePersonInTasks,
	}
	bus.RegisterSubscriptionsOnBus(subscriptionList)
}

func updatePersonInTasks(event microbus.DataEvent) {
	pers := event.Data.(*personpb.Person)
	//TODO go through all records which uses the person with pers.Id and update the display name
	fmt.Println("task mocker: simulate long running task")
	time.Sleep(15 * time.Second)
	fmt.Println("task mocker:", pers)

}
