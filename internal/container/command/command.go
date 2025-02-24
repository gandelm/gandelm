package command

import (
	"fmt"
	"os"
	"os/exec"
)

type Command struct {
}

func (c *Command) GitClone(org, repo string) error {
	return c.run(
		"git",
		"clone",
		fmt.Sprintf("https://github.com/%s/%s.git", org, repo),
		"./repositories/"+repo,
	)
}

func (c *Command) HelmInstall(releaseName, path string) error {
	return c.run(
		"helm",
		"install",
		"app",
		path,
		"--namespace",
		releaseName,
		"--create-namespace",
	)
}

func (c *Command) HelmUnInstall(releaseName string) error {
	return c.run(
		"helm",
		"uninstall",
		"app",
		"--namespace",
		releaseName,
	)
}

func (c *Command) run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
