package command

import (
    "github.com/spf13/cobra"
)

var env = &cobra.Command{
    Use: "env",
    Run: func(cmd *cobra.Command, args []string) {
    },
}

func init(){
    CMD.AddCommand(env)
}
