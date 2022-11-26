package main

import (
	"Api_Mux/api"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	//create api object
	a := &api.Api{}
	//register routes
	a.RegisterRoutes(r)

	r.HandleFunc("/", handleIndex).Methods(http.MethodGet) //se puedn definir los metodos de una vez, mejor forma es esta
	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	log.Println("Listening...")
	srv.ListenAndServe()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("{\"message\": \"Hello World\"}")
}
