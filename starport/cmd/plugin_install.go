package starportcmd

import (
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
	s := clispinner.New().SetText("Fetching Plugins...\n")
	defer s.Stop()

	c, err := newChainWithHomeFlags(cmd, chain.EnableThirdPartyModuleCodegen())
	if err != nil {
		return err
	}

	plugins, err := starplug.Init(c)
	if err != nil {
		return err
	}

	err = starplug.Install(cmd.Context(), plugins)
	return nil
}