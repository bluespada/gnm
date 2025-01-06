package command

import (
    "github.com/spf13/cobra"
)

var destroy = &cobra.Command{
    Use: "destroy",
    Aliases: []string{"d"},
}

func init(){
    CMD.AddCommand(destroy)
}
