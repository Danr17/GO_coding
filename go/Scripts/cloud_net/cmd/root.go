package cmd

import (
	"github.com/spf13/cobra"
)

var cloud string

func GetRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:               "cloud_net",
		Short:             "Cloud Networking Tool",
		SilenceUsage:      true,
		DisableAutoGenTag: true,
		Long: `Cloud Net is a multi-cloud troubleshooting and visualization networking tool`,
	}

	// rootCmd.SetArgs(args)

	rootCmd.PersistentFlags().StringVarP(&cloud, "cloud", "c", "", "Cloud vendor (required)")

	// cmd.AddFlags(rootCmd)

	rootCmd.AddCommand(discoverCmd())
	rootCmd.AddCommand(watchCmd())
	rootCmd.AddCommand(listCmd())

	return  rootCmd
}
