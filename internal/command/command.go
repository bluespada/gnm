package command

import (
    "github.com/spf13/cobra"
)


var CMD = &cobra.Command{
    Use: "gnm",
    Short: "Node Manager",
    Long: "Simple tools for manage multiple node js version.",
    RunE: func(cmd *cobra.Command, args []string) error {
        return nil
    },
}
