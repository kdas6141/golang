package main

import (
	"context"
    "net/http"
    "log"
	"os"
	"time"
	"os/signal"
	"github.com/gorilla/mux"
	"microservice/handlers"
    "microservice/env"
 )

var bindAddress = env.String("BIND_ADDRERSS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "sample-api", log.LstdFlags)

	//create the handlers
	ph := handlers.NewProducts(l)

	//create the new serve mux and register the handler
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

    putRouter := sm.Methods(http.MethodPut).Subrouter()
    putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)

    postRouter := sm.Methods(http.MethodPost).Subrouter()
    postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	s := &http.Server {
		Addr:	":9090",
		Handler:	sm,
		IdleTimeout: 120*time.Second,
        ReadTimeout: 1*time.Second,
        WriteTimeout: 1*time.Second,
	}

	go func() {
		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
    signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}

