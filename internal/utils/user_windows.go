//go:build windows

package utils

import (
	"os/user"
)

func GetRealUser() (*user.User, error) {
	return user.Current()
}

func DropPrivileges(u *user.User) error {
	// No-op on Windows — privilege dropping requires token manipulation
	// which is out of scope here
	return nil
}
