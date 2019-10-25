package commands

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "repo2graph is a powerful repository diagram generator",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Running Server!")
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
