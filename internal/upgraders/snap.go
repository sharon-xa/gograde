package upgraders

import (
	"context"
	"os/exec"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Snap struct {
	privileged bool
}

func NewSnap() *Snap {
	return &Snap{privileged: true}
}

func (s *Snap) Name() string {
	return "snap"
}

func (s *Snap) Privileged() bool {
	return s.privileged
}

func (s *Snap) Available() bool {
	_, err := exec.LookPath(s.Name())
	if err != nil {
		return false
	} else {
		return true
	}
}

func (s *Snap) Run(ctx context.Context) error {
	err := utils.RunCmd(ctx, s.Name(), "refresh")
	if err != nil {
		return err
	}
	return nil
}
