package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/peeb01/todo-app/internal/config"
	"github.com/peeb01/todo-app/internal/db"
	router "github.com/peeb01/todo-app/internal/route"
	"github.com/peeb01/todo-app/migration"

	"github.com/spf13/cobra"
)

var port string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start API server",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.Load()

		fmt.Println("Connecting to dabase")
		db.ConnectDatabase(cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME, cfg.DB_PORT, cfg.DB_SSLMODE, cfg.DB_TIMEZONE)
		migration.Migration(db.DB)
		fmt.Println("Database is connected")

		app := router.New(db.DB)
		app.HideBanner = false
		app.HidePort = false
		app.Logger.Fatal(app.Start(":" + cfg.Port))


		if cmd.Flags().Changed("port") {
			cfg.Port = port
		}

		srv := &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: app,
		}

		go func() {
			fmt.Println("Server running on port:", cfg.Port)
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %s\n", err)
			}
		}()

		/*
			Wait for an interrupt signal to gracefully shut down the server
			and accept graceful shutdowns from SIGINT (Ctrl+C) and SIGTERM (used by Docker, K8s)
		*/
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		// Block until we receive our signal
		<-quit
		log.Println("Shutting down server...")

		// Create a deadline to wait for
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("Server forced to shutdown: %v", err)
		}

		// Close the database connection
		sqlDB, err := db.DB.DB()
		if err != nil {
			log.Printf("Error getting underlying database for closing: %v", err)
		} else {
			log.Println("Closing database connection...")
			sqlDB.Close()
			log.Println("Database connection closed.")
		}

		log.Println("Server exiting")
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
