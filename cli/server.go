package cli

import (
	"fmt"
	"log"

	"github.com/drewart/hen/internal/server"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "server run a hen server",
	Long:  `server run a hen encryption service`,
	Run: func(cmd *cobra.Command, args []string) {
		port, _ := cmd.Flags().GetInt("port")
		runServer(port)
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntP("port", "p", 3000, "port number default 3000")
}

func runServer(port int) {
	err := server.Setup()
	if err != nil {
		log.Fatal(err)
	}

	r := server.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(fmt.Sprintf(":%d", port))
}
