package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Winget struct {
	privileged bool
}

func NewWinget() *Winget {
	return &Winget{privileged: false}
}

func (p *Winget) Name() string {
	return "winget"
}

func (p *Winget) Privileged() bool {
	return p.privileged
}

func (p *Winget) Available() bool {
	_, err := exec.LookPath(p.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (p *Winget) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, p.Name(), "upgrade", "--all", "--silent")
	if err != nil {
		return err
	}

	return nil
}
