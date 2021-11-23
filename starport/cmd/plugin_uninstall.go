package starportcmd

import "github.com/spf13/cobra"

func NewPluginUninstall() *cobra.Command {
	c := &cobra.Command{
		Use:   "uninstall",
		Short: "uninstall all plugins",
		Long:  `TODO FILL`,
		RunE:  pluginUninstallHandler,
	}

	return c
}

func pluginUninstallHandler(cmd *cobra.Command, args []string) error {
	return nil
}
