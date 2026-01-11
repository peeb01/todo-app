package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo-app",
	Short: "Todo App CLI",
	Long:  "Todo App CLI",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run command: go run main.go serve")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
