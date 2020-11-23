package eventHandler

import (
	"github.com/theNorstroem/todo-server/pkg/person/core"
	"github.com/theNorstroem/todo-server/pkg/person/interfaces/eventHandler"
	"github.com/veith/fgs-lib/pkg/microbus"
)

// Add your subcriptions to this list
func (s subscribtions) registerSubscriptions() {
	// DO NOT  make multiple funcs per topic possible
	// you can wrap the multiple functions in a function that calls the other ones, if needed
	var subscriptionList = map[string]func(event microbus.DataEvent){
		"example":                s.Example,
		"person.Persons.Updated": s.LogUpdatedPerson,
	}
	microbus.RegisterSubscriptionsOnBus(subscriptionList, s.eventBus)
}

// Implementations are in impl.go
// Add other needed services for your implementation here. Usually you should be ok with the domainservices of the
// current module. If you need more, you have to pass them to RegisterSubscriptions
type subscribtions struct {
	eventBus       *microbus.EventBus
	PersonServices *core.DomainServices
}

// RegisterSubscriptions register the topic subscriptions on the bus.
func RegisterSubscriptions(bus *microbus.EventBus, services *core.DomainServices) eventHandler.Subscribtions {
	s := &subscribtions{bus, services}
	// register the subscirbers
	s.registerSubscriptions()
	return s
}
