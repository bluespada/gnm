package command

import (
    "github.com/spf13/cobra"
)

var cswitch = &cobra.Command{
    Use: "switch",
    Aliases: []string{"s"},
}

func init(){
    CMD.AddCommand(cswitch)
}
