package extpoints

import (
	"golang.org/x/net/context"
)

type Connector interface {
	UID() string
	Connect(context.Context) (context.Context, error)
	Close(context.Context) error
}
