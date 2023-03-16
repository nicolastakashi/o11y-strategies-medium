package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/payments", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Hello Payments!")
	})

	mux.HandleFunc("/checkouts", func(w http.ResponseWriter, req *http.Request) {
		res, err := http.Get("http://proxy:8080/payments")
		if err != nil {
			fmt.Printf("error making http request: %s\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(w, "client: status code: %d\n", res.StatusCode)
	})

	httpServer := &http.Server{
		Addr:        ":8081",
		Handler:     mux,
		BaseContext: func(_ net.Listener) context.Context { return ctx },
	}
	// if your BaseContext is more complex you might want to use this instead of doing it manually
	// httpServer.RegisterOnShutdown(cancel)

	// Run server
	go func() {
		if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
			// it is fine to use Fatal here because it is not main gorutine
			log.Fatalf("HTTP server ListenAndServe: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)

	signal.Notify(
		signalChan,
		syscall.SIGHUP,  // kill -SIGHUP XXXX
		syscall.SIGINT,  // kill -SIGINT XXXX or Ctrl+c
		syscall.SIGQUIT, // kill -SIGQUIT XXXX
	)

	<-signalChan
	log.Print("os.Interrupt - shutting down...\n")

	go func() {
		<-signalChan
		log.Fatal("os.Kill - terminating...\n")
	}()

	gracefullCtx, cancelShutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelShutdown()

	if err := httpServer.Shutdown(gracefullCtx); err != nil {
		log.Printf("shutdown error: %v\n", err)
		defer os.Exit(1)
		return
	} else {
		log.Printf("gracefully stopped\n")
	}

	// manually cancel context if not using httpServer.RegisterOnShutdown(cancel)
	cancel()

	defer os.Exit(0)
	return
}
