package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Pacman struct {
	privileged bool
}

func NewPacman() *Pacman {
	return &Pacman{privileged: true}
}

func (p *Pacman) Name() string {
	return "pacman"
}

func (p *Pacman) Privileged() bool {
	return p.privileged
}

func (p *Pacman) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Pacman) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "-Syu", "--noconfirm")
	if err != nil {
		return err
	}
	return nil
}
