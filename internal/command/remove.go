package command

import (
    "github.com/spf13/cobra"
)

var remove = &cobra.Command{
    Use: "remove",
    Aliases: []string{"r"},
}

func init(){
    CMD.AddCommand(remove)
}
