// Package utils provides small cross-platform helpers for command execution,
// OS checks, and user/privilege operations used by CLI workflows.
package utils

import (
	"context"
	"os"
	"os/exec"
)

func RunCmd(ctx context.Context, cmdName string, arg ...string) error {
	cmd := exec.CommandContext(ctx, cmdName, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
