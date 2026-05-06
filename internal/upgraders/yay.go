package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Yay struct {
	privileged bool
}

func NewYay() *Yay {
	return &Yay{privileged: false}
}

func (p *Yay) Name() string {
	return "yay"
}

func (p *Yay) Privileged() bool {
	return p.privileged
}

func (p *Yay) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Yay) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "-Sua", "--noconfirm")
	if err != nil {
		return err
	}
	return nil
}
