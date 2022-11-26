package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type Api struct{}

//propiedad squema de mux es para maear los parametros de un struct en un equest
//go get github.com/gorilla/schema

//para que funcione el schema se debe de crear un struct con los campos que se quieren mapear

type BooksParams struct {
	Limit int `schema:"limit"` //la ventaja de esto es que se puede mandar el parametro como un key y agregar otros facilmente ya que el decoder los mapea por nosotros
}

type PostBook struct {
	Title string `json:"title"`
}

var (
	//decoder de nuestro schema
	decoder = schema.NewDecoder()
	books   = []string{"The Hitchhiker's Guide to the Galaxy", "The Restaurant at the End of the Universe", "Life, the Universe and Everything", "So Long, and Thanks for All the Fish", "Mostly Harmless"}
)

// parametro para controlar el tamanno de la respuesta
func (a *Api) getBooks(w http.ResponseWriter, r *http.Request) {

	params := &BooksParams{}

	err := decoder.Decode(params, r.URL.Query()) //nuestro aprams y nuestro source
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if params.Limit < 0 || params.Limit > len(books) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var to int

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}

	//desde el inicio hasta el limit
	json.NewEncoder(w).Encode(books[:to])
	//en el post man mandar parametro limit = num como un key
}

func (a *Api) getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //mapa de string

	idParam := params["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	index := id - 1

	if index < 0 || index >= len(books)-1 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(books[index]) //esto es manejo de parametros de un array no de un struct
}

func (a *Api) postBook(w http.ResponseWriter, r *http.Request) {
	book := &PostBook{}

	err := json.NewDecoder(r.Body).Decode(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	books = append(books, book.Title)
	w.WriteHeader(http.StatusCreated)
}
