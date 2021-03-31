package main

import (
	"os"

	"github.com/danrusei/cloud_net/cmd"
)


func main() {
	// if err := cmd.GetConfig(); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Could not initialize: %v\n", err)
	// 	os.Exit(1)
	// }
	rootCmd := cmd.GetRootCmd(os.Args[1:])

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}


