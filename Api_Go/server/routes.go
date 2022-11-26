package server

import (
	"fmt"
	"net/http"
)

// aqui se vana d efinir las rutas que van a llamar a los handlers
// El handleFunc va a ser nuestro manejador de rutas el cual recibe una ruta y una funcion que va a realizar cuando se ingrese
func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w, r)

		case http.MethodPost:
			addCountry(w, r)

		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
			return
		}
	})
}
