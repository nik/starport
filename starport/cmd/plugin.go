package starportcmd

import (
	"github.com/spf13/cobra"
)

func NewPlugin() *cobra.Command {
	c := &cobra.Command{
		Use:   "plugin [command]",
		Short: "Download, install, remove plugins",
		Long: `TODO: Fill out

		TODO
		`,
		Args: cobra.ExactArgs(1),
	}

	c.AddCommand(NewPluginInstall())
	c.AddCommand(NewPluginUpdate())
	c.AddCommand(NewPluginUninstall())

	return c
}
