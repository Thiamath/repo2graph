package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "repo2graph",
	Short: "repo2graph is a powerful repository diagram generator",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
