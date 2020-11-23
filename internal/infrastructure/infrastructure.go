package infrastructure

import (
	"github.com/upper/db/v4"
	"github.com/veith/fgs-lib/pkg/microbus"
	"google.golang.org/grpc"
)

type infrastructure struct {
	Grpc             *grpc.Server
	InternalEventBus *microbus.EventBus
	DbSession        db.Session
}

var Infrastructure infrastructure
