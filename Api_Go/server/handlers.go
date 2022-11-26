package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// aqui manejaremos todas las rutas

func index(w http.ResponseWriter, r *http.Request) {

	//CON ESTO NOS ASEGURAMOS DE QUE SE UTILICE UN METODO DE INGRESA CORRECTO
	if r.Method != http.MethodGet {
		//DEVUELVE UN ESTADO, LOS ESTADOS ESTAN EN LOS HEADERS
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method not allowed")
		return
	}

	//HACEMOS UN PRINT DE LOS DATOS CUANDO SE INGRESE
	fmt.Fprintf(w, "hello there %s", "visitor")

}

//endpoints

func getCountries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)
}

func addCountry(w http.ResponseWriter, r *http.Request) {
	country := &Country{}
	err := json.NewDecoder(r.Body).Decode(country)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "%v", err)
		return
	}

	countries = append(countries, country)
	fmt.Fprintf(w, "Country was addded")
}
