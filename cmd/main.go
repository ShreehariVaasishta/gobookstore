package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShreehariVaasishta/gobookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "gorm.io/driver/mysql"
)

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(handleNotFound)
	routes.RegisteredBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	errorResponse := map[string]string{"detail": "404 Not found"}
	jsonResponse, _ := json.Marshal(errorResponse)

	w.Write(jsonResponse)
}
