package upgraders

import (
	"context"
	"os/exec"
	"os/user"
	"strings"

	"github.com/sharon-xa/gograde/internal/utils"
)

type Npm struct {
	privileged bool
	executable string
	realUser   *user.User
}

func NewNpm() *Npm {
	if utils.IsWindows() {
		// no executable name in windows
		return &Npm{privileged: false}
	}

	if _, err := exec.LookPath("fnm"); err == nil {
		return &Npm{privileged: false, executable: "fnm"}
	}

	if prefix, err := exec.Command("brew", "--prefix").Output(); err == nil {
		if npmPath, err := exec.LookPath("npm"); err == nil {
			if strings.Contains(npmPath, strings.TrimSpace(string(prefix))) {
				return &Npm{privileged: false, executable: "npm"}
			}
		}
	}
	return &Npm{privileged: true}
}

func (n *Npm) Name() string {
	return "npm"
}

func (n *Npm) Privileged() bool {
	return n.privileged
}

func (n *Npm) Available() bool {
	switch n.executable {
	case "fnm":
		_, err := exec.LookPath("fnm")
		return err == nil
	// case "nvm":
	// 	nvmDir := filepath.Join(n.realUser.HomeDir, ".nvm")
	// 	_, err := os.Stat(nvmDir)
	// 	return err == nil
	default:
		_, err := exec.LookPath(n.Name())
		return err == nil
	}
}

func (n *Npm) Args() []string {
	switch n.executable {
	case "fnm":
		return []string{"fnm", "exec", "--using", "default", "--", "npm", "update", "-g"}
	// case "nvm":
	// 	nvmScript := filepath.Join(n.realUser.HomeDir, ".nvm", "nvm.sh")
	// 	return []string{
	// 		"bash",
	// 		"-c",
	// 		fmt.Sprintf("source %s && nvm exec default npm update -g", nvmScript),
	// 	}
	default:
		return []string{n.Name(), "update", "-g"}
	}
}

func (n *Npm) Run(ctx context.Context) error {
	args := n.Args()
	return utils.RunCmd(ctx, args[0], args[1:]...)
}
