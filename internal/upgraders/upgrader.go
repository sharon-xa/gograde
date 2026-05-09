// Package upgraders is where we implmenet the package manager upgrader,
// which will implement the interface below
package upgraders

import (
	"context"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Upgrader interface {
	Name() string
	Available() bool
	Privileged() bool
	Run(ctx context.Context) error
}

func GetUpgradersBasedOnOS() []Upgrader {
	if utils.IsWindows() {
		return []Upgrader{
			NewWinget(),
			NewScoop(),
			NewChoco(),

			NewNpm(),
		}
	}

	return []Upgrader{
		NewApt(),
		NewBrew(),
		NewDnf(),
		NewFlatpak(),
		NewPacman(),
		NewParu(),
		NewYay(),
		NewPort(),

		NewNpm(),
	}
}
