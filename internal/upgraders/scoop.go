package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Scoop struct {
	privileged bool
}

func NewScoop() *Scoop {
	return &Scoop{privileged: false}
}

func (s *Scoop) Name() string {
	return "scoop"
}

func (s *Scoop) Privileged() bool {
	return s.privileged
}

func (s *Scoop) Available() bool {
	_, err := exec.LookPath(s.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Scoop) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, s.Name(), "upgrade", "all", "-y")
	if err != nil {
		return err
	}

	return nil
}
