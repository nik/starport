package plugin

import (
	"context"
	"fmt"
	"os"
)

const pluginFolder = ".plugins"

// Install fetches all repos from GitHub assuming no dependencies and returns
// the installed list
func Install(ctx context.Context, plugins []plugin) ([]*string, error) {
	var installed []*string

	err := scaffoldPluginsDir()
	if err != nil {
		return nil, err
	}

	// TODO check if plugins already exist within folder
	for _, p := range plugins {
		err = p.Cmd.Create(installFolder(p.Name), p.RepoUrl)
		if err != nil {
			return nil, err
		}
		installed = append(installed, &p.Name)
	}

	return installed, nil
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