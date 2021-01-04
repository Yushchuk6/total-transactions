package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Yushchuk6/total-transactions/api/handler"
	"github.com/Yushchuk6/total-transactions/config"
	"github.com/Yushchuk6/total-transactions/entity"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	client := entity.NewClient(config.Scheme, config.Host, config.APIKey)
	getTotal := handler.GetTotalByID(*client)
	router.Handler("GET", "/api/block/:id/total", getTotal)

	s := http.Server{
		Addr:    ":8000",
		Handler: router,
	}

	go func() {
		log.Println("Starting server on port 8000")

		err := s.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	log.Println("Shutdonw signal: ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
