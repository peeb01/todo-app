package cmd

import (
	"fmt"
	"net/http"

	"github.com/peeb01/todo-app/internal/config"
	"github.com/peeb01/todo-app/internal/route"

	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Run: func(cmd *cobra.Command, args []string) {

		router := router.New()

		cfg := config.Load()

		if cmd.Flags().Changed("port") {
			cfg.Port = port
		}

		fmt.Println("Server running on port:", cfg.Port)
		err := http.ListenAndServe(":"+cfg.Port, router)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// --port override config
	serveCmd.Flags().StringVarP(&port, "port", "p", "", "server port")
}
