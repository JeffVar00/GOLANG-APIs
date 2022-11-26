package api

import (
	"github.com/labstack/echo/v4"
)

func (a *Api) RegisterRoutes(e *echo.Echo) {

	e.Use(requestIDHandler)

	public := e.Group("")                    //echo usa grupos en lugar de routers
	protected := e.Group("", authMiddleware) //a los grupos se le mandan los middlewares que queremos usar

	//handlers
	public.GET("/books", a.getBooks)
	public.GET("/books/:id", a.getBook) //el :id es un parametro ya no usa llaves como en mux
	protected.POST("/books", a.postBook)

}
