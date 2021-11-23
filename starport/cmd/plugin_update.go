package starportcmd

import "github.com/spf13/cobra"

func NewPluginUpdate() *cobra.Command {
	c := &cobra.Command{
		Use:   "update",
		Short: "updates all plugins",
		Long:  `TODO FILL`,
		RunE:  pluginUpdateHandler,
	}

	return c
}

func pluginUpdateHandler(cmd *cobra.Command, args []string) error {
	return nil
}
