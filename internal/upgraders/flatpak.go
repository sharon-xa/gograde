package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Flatpak struct {
	privileged bool
}

func NewFlatpak() *Flatpak {
	return &Flatpak{privileged: false}
}

func (p *Flatpak) Name() string {
	return "flatpak"
}

func (p *Flatpak) Privileged() bool {
	return p.privileged
}

func (p *Flatpak) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Flatpak) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "update", "-y")
	if err != nil {
		return err
	}
	return nil
}
