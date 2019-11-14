package commands

import (
	"github.com/Thiamath/repo2graph/pkg/server"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var debugLevel bool

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "repo2graph is a powerful repository diagram generator",
	Run: func(cmd *cobra.Command, args []string) {
		if debugLevel {
			log.SetLevel(log.DebugLevel)
			log.Debug("Activating DEBUG level")
		}
		log.Info("Running Server!")
		server.StartServer()
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&debugLevel, "debug", false, "Set debug level")

	rootCmd.AddCommand(serverCmd)
}
