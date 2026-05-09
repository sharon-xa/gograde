package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Choco struct {
	privileged bool
}

func NewChoco() *Choco {
	return &Choco{privileged: false}
}

func (c *Choco) Name() string {
	return "choco"
}

func (c *Choco) Privileged() bool {
	return c.privileged
}

func (c *Choco) Available() bool {
	_, err := exec.LookPath(c.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (c *Choco) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, c.Name(), "upgrade", "all", "-y")
	if err != nil {
		return err
	}

	return nil
}
