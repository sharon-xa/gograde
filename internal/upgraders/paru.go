package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Paru struct {
	privileged bool
}

func NewParu() *Paru {
	return &Paru{privileged: false}
}

func (p *Paru) Name() string {
	return "paru"
}

func (p *Paru) Privileged() bool {
	return p.privileged
}

func (p *Paru) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Paru) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "-Sua", "--noconfirm", "--skipreview")
	if err != nil {
		return err
	}
	return nil
}
