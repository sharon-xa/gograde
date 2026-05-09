package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Port struct {
	privileged bool
}

func NewPort() *Port {
	return &Port{privileged: true}
}

func (p *Port) Name() string {
	return "port"
}

func (p *Port) Privileged() bool {
	return p.privileged
}

func (p *Port) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Port) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "selfupdate")
	if err != nil {
		return err
	}

	err = utils.RunCmd(ctx, p.Name(), "upgrade", "outdated")
	if err != nil {
		return err
	}

	return nil
}
