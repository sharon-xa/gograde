package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Apt struct {
	privileged bool
}

func NewApt() *Apt {
	return &Apt{privileged: true}
}

func (p *Apt) Name() string {
	return "apt"
}

func (p *Apt) Privileged() bool {
	return p.privileged
}

func (p *Apt) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Apt) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "update")
	if err != nil {
		return err
	}

	err = utils.RunCmd(ctx, p.Name(), "full-upgrade", "-y")
	if err != nil {
		return err
	}
	return nil
}
