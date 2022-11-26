package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Api) RegisterRoutes(r *mux.Router) {

	//middlewares
	r.Use(requestIDHandler) //este middleware es general

	//ahora este lo que hace es generar un header personalizado con un id unico para cada request
	//para probarlo se puede usar postman y mandar un request, despues de eso se puede mandar otro request y ver que el id es diferente
	//para ver el id se puede ir a la pestaña de headers y ver el valor de X-Request-Id

	//usar middleware solo para algunas rutas
	public := r.NewRoute().Subrouter() //rama del router
	protected := r.NewRoute().Subrouter()

	//handlers
	public.HandleFunc("/books", a.getBooks).Methods(http.MethodGet)
	public.HandleFunc("/books/{id}", a.getBook).Methods(http.MethodGet)

	//solo quiero proteger un endpoint
	protected.HandleFunc("/books", a.postBook).Methods(http.MethodPost)

	//Para el authorized verifica si el header recibe una autorizacion
	protected.Use(authMiddleware)

}
