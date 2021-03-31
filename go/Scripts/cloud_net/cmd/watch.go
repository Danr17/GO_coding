package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func watchCmd() *cobra.Command {
	watchCmd := &cobra.Command{
		Use:     "watch",
		Aliases: []string{"w"},
		Short:   "Watch cloud networking resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			// cmd.HelpFunc()(cmd, args)
			if len(args) != 0 {
				return fmt.Errorf("unknown resource type %q", args[0])
			}

			return nil
		},
	}

	return watchCmd
}
