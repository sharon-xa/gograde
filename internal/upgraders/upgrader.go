// Package upgraders is where we implmenet the package manager upgrader,
// which will implement the interface below
package upgraders

import "context"

type Upgrader interface {
	Name() string
	Available() bool
	Privileged() bool
	Run(ctx context.Context) error
}
