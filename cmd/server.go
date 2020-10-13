package cmd

import (
	"github.com/Bnei-Baruch/study-material/api"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Serve the backend API",
	Run:   serverFn,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}

func serverFn(cmd *cobra.Command, args []string) {
	api.Init()
}
