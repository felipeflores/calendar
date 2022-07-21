package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Config is the http server config
type Config struct {
	Address         string
	WriteTimeout    time.Duration
	ReadTimeout     time.Duration
	IdleTimeout     time.Duration
	ShutdownTimeout time.Duration
}

// Run dispatch a goroutine with ListenAndServe() and will wait for
// a syscall to stop gracefully the http server.
func Run(c Config, middleware http.Handler) {

	if c.WriteTimeout < time.Second {
		c.WriteTimeout = time.Second * 15
	}

	if c.ReadTimeout < time.Second {
		c.ReadTimeout = time.Second * 15
	}

	if c.IdleTimeout < time.Second {
		c.IdleTimeout = time.Second * 60
	}

	if c.ShutdownTimeout < time.Second {
		c.ShutdownTimeout = time.Second * 10
	}

	srv := &http.Server{
		Addr: c.Address,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: c.WriteTimeout,
		ReadTimeout:  c.ReadTimeout,
		IdleTimeout:  c.IdleTimeout,
		Handler:      middleware,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			if err != http.ErrServerClosed {
				// Error starting or closing listener:
				log.Fatalf("HTTP server ListenAndServe: %v", err)
			}
		}
	}()

	ch := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-ch

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), c.ShutdownTimeout)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := srv.Shutdown(ctx); err != nil {
		// Error from closing listeners, or context timeout:
		log.Printf("HTTP server Shutdown: %v", err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Println("HTTP server stopped.")
}
