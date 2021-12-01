package starportcmd

import (
	"fmt"
	"github.com/tendermint/starport/starport/pkg/clispinner"
	"github.com/tendermint/starport/starport/services/chain"
	starplug "github.com/tendermint/starport/starport/services/plugin"

	"github.com/spf13/cobra"
)

func NewPluginInstall() *cobra.Command {
	c := &cobra.Command{
		Use:   "install",
		Short: "install all plugins",
		Long:  `TODO FILL`,
		RunE:  pluginInstallHandler,
	}

	return c
}

func pluginInstallHandler(cmd *cobra.Command, args []string) error {
	s := clispinner.New().SetText("Fetching Plugins...")
	defer s.Stop()

	c, err := newChainWithHomeFlags(cmd, chain.EnableThirdPartyModuleCodegen())
	if err != nil {
		return err
	}

	plugins, err := starplug.Init(c)
	if err != nil {
		return err
	}

	installed, err := starplug.Install(cmd.Context(), plugins)
	if err != nil {
		return err
	}

	s.Stop()

	for _, install := range installed {
		fmt.Printf("\n🎉 %s added. \n\n", *install)
	}

	starplug.LoadPlugins()
	loadedPlugins := starplug.GetLoadedPlugins()
	for _, lp := range loadedPlugins {
		fmt.Printf("manifest: %+v", lp.Manifest)
	}
	return nil
}