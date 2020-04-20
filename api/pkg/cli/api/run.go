package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/mux"
	"github.com/replicatedhq/kots-bots/api/pkg/handlers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		PreRun: func(cmd *cobra.Command, args []string) {
			viper.BindPFlags(cmd.Flags())
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			r := mux.NewRouter()
			r.Use(mux.CORSMethodMiddleware(r))

			r.HandleFunc("/healthz", handlers.Healthz).Methods("GET")

			// this is only served in kots
			if os.Getenv("DISABLE_SPA_SERVING") != "1" {
				fmt.Printf("serving static handler\n")
				spa := handlers.SPAHandler{StaticPath: filepath.Join("web", "dist"), IndexPath: "index.html"}
				r.PathPrefix("/").Handler(spa)
			} else {
				r.HandleFunc("/", handlers.Root)
			}

			srv := &http.Server{
				Handler:      r,
				Addr:         ":3000",
				WriteTimeout: 15 * time.Second,
				ReadTimeout:  15 * time.Second,
			}

			fmt.Printf("Starting kotsbots-server API on port %d...\n", 3000)

			log.Fatal(srv.ListenAndServe())

			return nil
		},
	}

	return cmd
}
