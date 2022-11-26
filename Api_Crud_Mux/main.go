package main

import (
	"crud_GO/handlers"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//con el strict slash lo que le decimos es que si o si necesita hacer match con al ruta impuesta en el codigo
	router := mux.NewRouter().StrictSlash(true)

	// Index Routes Esta es la ruta inicial del API
	router.HandleFunc("/", handlers.IndexRoute)

	// Tasks Routes, aqui definimos todas las rutas que traermos de los handlers
	router.HandleFunc("/tasks", handlers.CreateTask).Methods("POST") //con esto le podemos indicar de una el metodo
	router.HandleFunc("/tasks", handlers.GetTasks).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.GetOneTask).Methods("GET")
	router.HandleFunc("/tasks/{id}", handlers.DeleteTask).Methods("DELETE")
	router.HandleFunc("/tasks/{id}", handlers.UpdateTask).Methods("PUT")

	fmt.Println("Server started on port ", 3000)

	//con el log manejamos el error de una vez
	log.Fatal(http.ListenAndServe(":3000", router))
}
