package plugin

import (
	"context"
	"fmt"
	"os"
)

// Install fetches all repos from GitHub assuming no dependencies
func Install(ctx context.Context, pluginPaths []plugin) error {
	err := createPluginsDir()
	if err != nil {
		return err
	}

	// create `.plugins/` folder
	// cd into it
	// run local git clone on each repo
	// cache perhaps
	// run go build -buildmode=plugin within repo
	// extract SO file
	cleanup(ctx)
	return nil
}

func createPluginsDir() error {
	err := os.Mkdir(".plugins", 0755)
	if err != nil {
		return err
	}

	return nil
}

func cleanup(ctx context.Context) {
	fmt.Println("removing .plugins/")
	os.RemoveAll(".plugins")
}