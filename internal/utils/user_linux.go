//go:build linux

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"golang.org/x/sys/unix"
)

func getUserPath(u *user.User) string {
	cmd := exec.Command("sudo", "-u", u.Username, "-i", "sh", "-c", "echo $PATH")
	out, err := cmd.Output()
	if err != nil {
		// fallback to a sane default
		return strings.Join([]string{
			u.HomeDir + "/.local/bin",
			"/usr/local/bin",
			"/usr/bin",
			"/bin",
		}, ":")
	}
	return strings.TrimSpace(string(out))
}

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

	if err := unix.Setresgid(gid, gid, gid); err != nil {
		return fmt.Errorf("setresgid: %w", err)
	}

	if err := unix.Setresuid(uid, uid, uid); err != nil {
		return fmt.Errorf("setresuid: %w", err)
	}

	if unix.Getuid() != uid || unix.Geteuid() != uid {
		return fmt.Errorf("privilege drop failed: uid=%d euid=%d, expected %d",
			unix.Getuid(), unix.Geteuid(), uid)
	}

	if err := os.Setenv("HOME", u.HomeDir); err != nil {
		return fmt.Errorf("setenv HOME: %w", err)
	}

	if err := os.Setenv("USER", u.Username); err != nil {
		return fmt.Errorf("setenv USER: %w", err)
	}

	if err := os.Setenv("LOGNAME", u.Username); err != nil {
		return fmt.Errorf("setenv LOGNAME: %w", err)
	}

	path := getUserPath(u)
	if err := os.Setenv("PATH", path); err != nil {
		return fmt.Errorf("setenv PATH: %w", err)
	}

	return nil
}
