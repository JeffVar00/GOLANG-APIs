package main

import (
	"Api_Echo/api"
	"net/http"

	"github.com/labstack/echo/v4" //vigilar traer esto la v porque VSCODE trae la base
	//"github.com/gorilla/mux"
	//ahora usaremos echo
	//go get github.com/labstack/echo/v4
)

func main() {
	e := echo.New()
	//mapeando a una ruta
	e.GET("/", handleIndex)

	a := &api.Api{}
	a.RegisterRoutes(e)

	//como el listenandserve, directamente lo trata como un handler de error
	e.Logger.Fatal(e.Start(":8080"))
}

func handleIndex(e echo.Context) error { //ya no recibe los request y response, sino que recibe un contexto, y siempre deuvelve un error
	return e.JSON(http.StatusOK, map[string]string{"message": "Hello World"}) //devuelve un json
}
