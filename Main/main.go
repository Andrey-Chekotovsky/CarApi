package main

import (
	"fmt"
	"handlers/cars"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("lets go")
	l := log.New(os.Stdout, "products-api ", log.LstdFlags)
	sm := mux.NewRouter()

	c := cars.Cars{L: l}

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/products", c.GetAll)
	getR.HandleFunc("/products/{id:[0-9]+}", c.Get)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/products", c.AddCar)
	postR.Use(c.MiddlewareValidateCar)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/products", c.Put)
	putR.Use(c.MiddlewareValidateCar)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/products/{id:[0-9]+}", c.Delete)

	server := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  100 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	server.ListenAndServe()
}
