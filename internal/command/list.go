package command

import (
    "github.com/spf13/cobra"
    "github.com/bluespada/gnm/internal/node"
)

var list = &cobra.Command{
    Use: "list",
    Aliases: []string{"ls"},
    Run: func(cmd *cobra.Command, args []string) {
        node.ListInstalled()
    },
}

func init(){
    CMD.AddCommand(list)
}
