package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Zypper struct {
	privileged bool
}

func NewZypper() *Zypper {
	return &Zypper{privileged: true}
}

func (z *Zypper) Name() string {
	return "zypper"
}

func (z *Zypper) Privileged() bool {
	return z.privileged
}

func (z *Zypper) Available() bool {
	_, err := exec.LookPath(z.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (z *Zypper) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, z.Name(), "update", "-y")
	if err != nil {
		return err
	}
	return nil
}
