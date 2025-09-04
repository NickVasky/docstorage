package main

import (
	"context"
	"log"

	"github.com/NickVasky/docstorage/internal/app"
)

func main() {
	//rootCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	//defer stop()

	// cfg, err := config.GetConfig()
	// if err != nil {
	// 	log.Fatalf("Error loading App config: %v", err)
	// }

	ctx := context.Background()

	a, err := app.NewApp(ctx)
	if err != nil {
		log.Fatalf("Failed to init app: %v", err)
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("Failed to run app: %v", err)
	}

	// go func() {
	// 	log.Printf("Starting server on %v...", server.Addr)
	// 	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
	// 		log.Fatalf("HTTP server error: %v", err)
	// 	}
	// 	log.Println("Stopped serving new connections.")
	// }()

	// <-rootCtx.Done()
	//Shutdown(server, db)
}
