package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/zYasser/MyFitness/controller"
	"github.com/zYasser/MyFitness/utils"
)

type Connector struct {
	router *mux.Router
}

var bindAddress = ":9090"

func main() {
	l := utils.GetLogger()

	con := &Connector{}
	con.router = controller.InitApplication().Router
	s := http.Server{
		Addr:         bindAddress,       // configure the bind address
		Handler:      con.router,        // set the default handler
		ErrorLog:     l.ErrorLog,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.InfoLog.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.ErrorLog.Fatalf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	l.InfoLog.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
}
