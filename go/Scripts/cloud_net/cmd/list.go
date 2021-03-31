package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func listCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Short:   "List cloud networking resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			// cmd.HelpFunc()(cmd, args)
			if len(args) != 0 {
				return fmt.Errorf("unknown resource type %q", args[0])
			}

			return nil
		},
	}

	return listCmd
}
