package api

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

//paquete go get -u github.com/google/uuid para generar ids

//to do esto debe retornar un handler

//ventasjas: podemos concatenar middlewares, podemos hacer un middleware que se encargue de la autenticacion, otro que se encargue de la autorizacion, otro que se encargue de la validacion de los datos, etc

var users = map[string]string{"user1": "password1", "user2": "password2"}

func requestIDHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		//obtener el id de la request
		requestID := e.Request().Header.Get("X-REQUEST-ID") //se acostumbra el prefijo de X cuando usamos headers custom

		//si esto no se ejecuta el valor que se mantiene el que recibe
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		e.Response().Header().Set("X-REQUEST-ID", requestID)

		return next(e)
	}
}

// mapa de valores para todos los usuarios que pueden ingresar a cada endpoint
func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {

		user := e.Request().Header.Get("Authorization")

		if users[user] == "" {
			return e.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(e)
	}
}
