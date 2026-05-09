//go:build darwin

package runner

import (
	"context"
	"fmt"
	"log"

	"github.com/sharon-xa/gograde/internal/upgraders"
	"github.com/sharon-xa/gograde/internal/utils"
)

func Run(ctx context.Context, upgraders []upgraders.Upgrader) []Result {
	var results []Result

	fmt.Println("")
	fmt.Println("=== Upgrading Privileged Package Managers ===")
	for _, u := range upgraders {
		if !u.Available() || !u.Privileged() {
			continue
		}

		fmt.Printf("\n")
		fmt.Printf("==> Upgrading %s\n", u.Name())
		err := u.Run(ctx)
		results = append(results, Result{Name: u.Name(), Err: err})
	}

	realUser, err := utils.GetRealUser()
	if err != nil {
		log.Fatalln(err)
	}

	err = utils.DropPrivileges(realUser)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("")
	fmt.Println("=== Upgrading Other Package Managers ===")
	for _, u := range upgraders {
		if !u.Available() || u.Privileged() {
			continue
		}

		fmt.Printf("\n")
		fmt.Printf("==> Upgrading %s\n", u.Name())
		err := u.Run(ctx)
		results = append(results, Result{Name: u.Name(), Err: err})
	}

	return results
}
