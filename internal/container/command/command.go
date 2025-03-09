package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/gandelm/gandelm/internal/container"
)

var _ container.Commander = (*Command)(nil)

type Command struct {
	ManifestPath string
	PackageName  string
}

func NewCommand(manifestPath, packageName string) *Command {
	return &Command{
		ManifestPath: manifestPath,
		PackageName:  packageName,
	}
}

func (c *Command) Manifest() string {
	return fmt.Sprintf("./%s/%s", c.ManifestPath, c.PackageName)
}

func (c *Command) GitClone(org, repo string) error {
	return c.run("git", "clone", fmt.Sprintf("https://github.com/%s/%s.git", org, repo), "./repositories/"+repo)
}

func (c *Command) HelmInstall(namespace, releaseName string) error {
	return c.run("helm", "install", releaseName, c.Manifest(), "--namespace", namespace, "--create-namespace")
}

func (c *Command) HelmUpgrade(namespace, releaseName string) error {
	return c.run("helm", "upgrade", releaseName, c.Manifest(), "--namespace", namespace)
}

func (c *Command) HelmUnInstall(namespace, releaseName string) error {
	return c.run("helm", "uninstall", releaseName, "--namespace", namespace)
}

func (c *Command) run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
