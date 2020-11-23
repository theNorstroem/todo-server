package eventHandler

import (
	"fmt"
	"github.com/veith/fgs-lib/pkg/microbus"
)

func (s subscribtions) Example(event microbus.DataEvent) {
	fmt.Println(event)
}

func (s subscribtions) LogUpdatedPerson(event microbus.DataEvent) {
	//fmt.Println(event)
}
