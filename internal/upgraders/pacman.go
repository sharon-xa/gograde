package upgraders

import "context"

type Pacman struct{}

func (p *Pacman) Name() string {
	return ""
}

func (p *Pacman) Available() bool {
	return false
}

func (p *Pacman) Run(ctx context.Context) error {
	return nil
}
