package runner

import (
	"context"
	"fmt"

	"github.com/sharon-xa/gograde/internal/upgraders"
)

type Result struct {
	Name string
	Err  error
}

func Run(ctx context.Context, upgraders []upgraders.Upgrader) []Result {
	var results []Result

	for _, u := range upgraders {
		if !u.Available() {
			continue
		}

		fmt.Printf("==> Upgrading %s\n", u.Name())
		err := u.Run(ctx)
		results = append(results, Result{Name: u.Name(), Err: err})
	}

	return results
}
