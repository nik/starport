package starportcmd

import "github.com/spf13/cobra"

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
	return nil
}
