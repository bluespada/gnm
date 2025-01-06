package command

import (
    "github.com/bluespada/gnm/internal/node"
    "github.com/spf13/cobra"
)

var install = &cobra.Command{
    Use: "install",
    Aliases: []string{"i"},
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        node.DownloadAndInstall(args[0])
    },
}

func init(){
    CMD.AddCommand(install)
}
