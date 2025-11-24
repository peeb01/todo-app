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
		fmt.Println("Hello From TODO App")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
