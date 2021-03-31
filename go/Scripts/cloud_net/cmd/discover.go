package cmd

import (
	"fmt"
	"os"

	"github.com/danrusei/cloud_net/pkg/gcp"
	"github.com/spf13/cobra"
)

var projectID string

func discoverCmd() *cobra.Command {
	discoverCmd := &cobra.Command{
		Use:     "discover",
		Aliases: []string{"dis"},
		Short:   "Discover cloud networking resources",
		RunE: func(cmd *cobra.Command, args []string) error {
			// cmd.HelpFunc()(cmd, args)
			// if len(args) != 0 {
			// 	return fmt.Errorf("unknown resource type %q", args[0])
			// }
			if projectID == "" {
				fmt.Fprintln(cmd.OutOrStderr(), "You need to provide a project. Use -p flag.")
				os.Exit(1)
			}
			if  len(os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")) == 0 {
				fmt.Fprintln(cmd.OutOrStderr(), "You need the credential setup, GOOGLE_APPLICATION_CREDENTIALS")
				os.Exit(1)
			}
			gcp.Discover(projectID)
			return nil
		},
	}

	discoverCmd.PersistentFlags().StringVarP(&projectID, "project", "p", "", "The project to be discovered")

//	discoverCmd.AddCommand(awsCmd())
//	discoverCmd.AddCommand(gcpCmd())
	return discoverCmd
}
