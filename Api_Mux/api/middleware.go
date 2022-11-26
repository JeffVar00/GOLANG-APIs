package api

import (
	"net/http"

	"github.com/google/uuid"
)

//paquete go get -u github.com/google/uuid para generar ids

//to do esto debe retornar un handler

//ventasjas: podemos concatenar middlewares, podemos hacer un middleware que se encargue de la autenticacion, otro que se encargue de la autorizacion, otro que se encargue de la validacion de los datos, etc

var users = map[string]string{"user1": "password1", "user2": "password2"}

func requestIDHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//obtener el id de la request
		requestID := r.Header.Get("X-Request-Id") //se acostumbra el prefijo de X cuando usamos headers custom

		//si esto no se ejecuta el valor que se mantiene el que recibe
		if len(requestID) == 0 {
			requestID = uuid.New().String()
		}

		w.Header().Set("X-Request-Id", requestID)

		next.ServeHTTP(w, r)
	})
}

// mapa de valores para todos los usuarios que pueden ingresar a cada endpoint
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		user := r.Header.Get("Authorization")

		if users[user] == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
