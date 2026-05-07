package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Brew struct {
	privileged bool
}

func NewBrew() *Brew {
	return &Brew{privileged: false}
}

func (p *Brew) Name() string {
	return "brew"
}

func (p *Brew) Privileged() bool {
	return p.privileged
}

func (p *Brew) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Brew) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "update")
	if err != nil {
		return err
	}

	err = utils.RunCmd(ctx, p.Name(), "upgrade", "-y")
	if err != nil {
		return err
	}

	err = utils.RunCmd(ctx, p.Name(), "upgrade", "--cask", "-y")
	if err != nil {
		return err
	}

	return nil
}
