package upgraders

import "context"

type Upgrader interface {
	Name() string
	Available() bool
	Privileged() bool
	Run(ctx context.Context) error
}
