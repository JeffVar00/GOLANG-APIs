package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Api struct{}

type BooksParams struct {
	Limit int `query:"limit"` //la ventaja de esto es que se puede mandar el parametro como un key y agregar otros facilmente ya que el decoder los mapea por nosotros
	//diferente de mux aqui se usa query
}

type PostBook struct {
	Title string `json:"title"`
}
type BookIdParams struct {
	Id int `param:"id"`
}

var (
	books = []string{"The Hitchhiker's Guide to the Galaxy", "The Restaurant at the End of the Universe", "Life, the Universe and Everything", "So Long, and Thanks for All the Fish", "Mostly Harmless"}
)

func (a *Api) getBooks(e echo.Context) error {

	params := &BooksParams{}

	err := e.Bind(params) //nuestro params y nuestro source, enlace de los aprametros de e con los del struct
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid Query Params")
	}

	if params.Limit < 0 || params.Limit > len(books) {
		return e.JSON(http.StatusBadRequest, "Invalid Query Params")
	}

	var to int

	if params.Limit > 0 {
		to = params.Limit
	} else {
		to = len(books)
	}

	//desde el inicio hasta el limit
	return e.JSON(http.StatusOK, books[:to])
}

func (a *Api) getBook(e echo.Context) error {

	params := &BookIdParams{} //echo maneja esto con el Param

	err := e.Bind(params)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid ID")
	}

	index := params.Id - 1

	if index < 0 || index >= len(books)-1 {
		return e.JSON(http.StatusBadRequest, "Invalid ID")
	}

	return e.JSON(http.StatusOK, books[index]) //esto es manejo de parametros de un array no de un struct
}

func (a *Api) postBook(e echo.Context) error {
	book := &PostBook{}

	err := e.Bind(book)
	if err != nil {
		return e.JSON(http.StatusBadRequest, "Invalid parameters")
	}

	books = append(books, book.Title)
	return e.NoContent(http.StatusCreated)
}
