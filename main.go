package main

import (
	"context"
	"controllers_example/controllers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ap := &controllers.ApplicationController{}
	err := controllers.NewController("/api", ap)
	if err != nil {
		panic(err)
	}

	mux := controllers.GetMux()

	s := &http.Server{
		Addr:         ":8000",
		Handler:      mux,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
