package command

import (
    "github.com/spf13/cobra"
)

var list = &cobra.Command{
    Use: "list",
    Aliases: []string{"ls"},
}

func init(){
    CMD.AddCommand(list)
}
