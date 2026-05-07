package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Dnf struct {
	privileged bool
}

func NewDnf() *Dnf {
	return &Dnf{privileged: true}
}

func (p *Dnf) Name() string {
	return "dnf"
}

func (p *Dnf) Privileged() bool {
	return p.privileged
}

func (p *Dnf) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Dnf) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "upgrade", "--refresh", "-y")
	if err != nil {
		return err
	}

	return nil
}
