package upgraders

import "context"

type Upgrader interface {
	Name() string
	Available() bool
	Run(ctx context.Context) error
}
