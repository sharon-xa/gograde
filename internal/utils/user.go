package utils

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
	"syscall"
)

func GetRealUser() (*user.User, error) {
	sudoUser := os.Getenv("SUDO_USER")
	if sudoUser == "" {
		return nil, fmt.Errorf("SUDO_USER not set — are you running with sudo?")
	}
	return user.Lookup(sudoUser)
}

func DropPrivileges(u *user.User) error {
	uid, err := strconv.Atoi(u.Uid)
	if err != nil {
		return fmt.Errorf("invalid uid %q: %w", u.Uid, err)
	}
	gid, err := strconv.Atoi(u.Gid)
	if err != nil {
		return fmt.Errorf("invalid gid %q: %w", u.Gid, err)
	}

	// GID must be dropped before UID — once you drop UID you can't change GID
	if err := syscall.Setgid(gid); err != nil {
		return fmt.Errorf("setgid: %w", err)
	}
	if err := syscall.Setuid(uid); err != nil {
		return fmt.Errorf("setuid: %w", err)
	}

	// Patch HOME so user-space tools find the right config/cache
	if err := os.Setenv("HOME", u.HomeDir); err != nil {
		return fmt.Errorf("setenv HOME: %w", err)
	}

	return nil
}
