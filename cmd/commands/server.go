package commands

import (
	"github.com/Thiamath/repo2graph/internal"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "repo2graph is a powerful repository diagram generator",
	Run: func(cmd *cobra.Command, args []string) {
		log.Info().Msg("Running Server!")
		internal.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
