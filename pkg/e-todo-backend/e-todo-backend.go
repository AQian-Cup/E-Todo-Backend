package e_todo_backend

import (
	"context"
	"e-todo-backend/pkg/db"
	"e-todo-backend/pkg/router"
	"errors"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	cfgFile string
)

func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:          "e-todo-backend",
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
	cmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "The path to the blog configuration file. Empty string for no configuration file.")
	cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	return cmd
}

func run() error {
	DBConfig := &db.PostgreSQLOptions{
		Host:     "43.133.185.245",
		User:     "e-todo-admin",
		Password: "e-todo-admin",
		Database: "e-todo",
		Port:     5432,
	}
	if err := DBConfig.Init(); err != nil {
		return err
	}
	r := gin.Default()
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:5173"}
	r.Use(cors.New(corsConfig))
	if err := router.InitRoutes(r); err != nil {
		return err
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown:", err)
	}
	log.Println("Server exiting")
	return nil
}
