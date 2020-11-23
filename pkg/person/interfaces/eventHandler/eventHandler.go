package eventHandler

import "github.com/veith/fgs-lib/pkg/microbus"

// Define methods that must be implemented by the eventHandler
type Subscribtions interface {
	LogUpdatedPerson(event microbus.DataEvent)
}
