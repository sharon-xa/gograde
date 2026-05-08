//go:build darwin

package utils

import (
	"fmt"
	"os"
	"os/user"
	"strconv"

	"golang.org/x/sys/unix"
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

	if err := unix.Setgroups([]int{gid}); err != nil {
		return fmt.Errorf("setgroups: %w", err)
	}

	if err := unix.Setgid(gid); err != nil {
		return fmt.Errorf("setgid: %w", err)
	}

	if err := unix.Setuid(uid); err != nil {
		return fmt.Errorf("setuid: %w", err)
	}

	if unix.Getuid() != uid || unix.Geteuid() != uid {
		return fmt.Errorf("privilege drop failed: uid=%d euid=%d, expected %d",
			unix.Getuid(), unix.Geteuid(), uid)
	}

	if err := os.Setenv("HOME", u.HomeDir); err != nil {
		return fmt.Errorf("setenv HOME: %w", err)
	}
	return nil
}
