package plugin

import (
	"context"
	"fmt"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/exec"
	"os"
)

const pluginFolder = ".plugins"

// Install fetches all repos from GitHub assuming no dependencies and returns
// the installed list
func Install(ctx context.Context, plugins []plugin) ([]*string, error) {
	var installed []*plugin
	var built []*string

	err := scaffoldPluginsDir()
	if err != nil {
		return nil, err
	}

	// TODO check if plugins already exist within folder
	for _, p := range plugins {
		location := installFolder(p.Name)
		err = p.Cmd.Create(location, p.RepoUrl)
		if err != nil {
			return nil, err
		}
		p.InstalledLocation = location
		installed = append(installed, &p)
	}

	for _, p := range installed {
		err = exec.Exec(ctx, []string{"make", "-C", fmt.Sprintf(".plugins/%s", p.Name), "build"})
		if err != nil {
			return []*string{}, err
		}
		built = append(built, &p.Name)
	}

	return built, nil
}

func scaffoldPluginsDir() error {
	if _, err := os.Stat(pluginFolder); os.IsNotExist(err) {
		err = os.Mkdir(pluginFolder, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func cleanup(ctx context.Context) {
	fmt.Println("removing .plugins/")
	os.RemoveAll(pluginFolder)
}

func installFolder(repo string) string {
	return fmt.Sprintf("%s/%s", pluginFolder, repo)
}