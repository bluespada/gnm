package command

import (
    "github.com/spf13/cobra"
    "github.com/bluespada/gnm/internal/node"
)

var listRemote = &cobra.Command{
    Use: "list-remote",
    Aliases: []string{"ls-remote"},
    Run: func(cmd *cobra.Command, args []string) {
        node.ListRemote()
    },
}

func init(){
    CMD.AddCommand(listRemote)
}
