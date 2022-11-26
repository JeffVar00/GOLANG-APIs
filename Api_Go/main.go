package main

import (
	"Api_Go/server"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	//contextos
	ctx := context.Background()

	//genera un channel que va a recebir senales por aprte del sistema operativo, y recibe solo una senal
	serverDoneChan := make(chan os.Signal, 1)

	//paquete signal. la propiedad notify va a enviar que reciba del so hacia uno de neustros channels, recibe a el como
	//parametro y se le ennvia la lista de senales a escuchar
	signal.Notify(serverDoneChan, os.Interrupt, syscall.SIGTERM)

	//AQUI FORMATEAMOS NUESTRO SERVIDOR EN EL PUERTO DESEADO
	srv := server.New(":8080")

	//go routine y go function
	go func() {
		err := srv.ListenAndServe() //ESTO NOS DEVUELVE UN ESTADO DEL SERVIDOR
		if err != nil {
			panic(err)
		}
	}()

	//ahora mostrar mensajes de salida
	log.Println("Server Started")

	//espera una senal, no es necesario que este en una variable
	<-serverDoneChan

	//al recibir una de las senales dadas entonces paso a lo que sigue en codigo
	err := srv.Shutdown(ctx)

	if err != nil {
		panic(err)
	}

	log.Println("Server Stoped")

}

//HAY QUE TENER EN CUENTA LOS CODIGOS DE ESTADO QUE DEVUELVE EL NAVEGADOR CUANDO OCURREN DIFERENTES REQUEST
//200 SIGNIFICA QUE TODO ESTA OK POR EJE
//ESTO LO PODEMOS ENCONTRAR EN LA DOCUMENTACION DE MOZILLA
//ES NECESARIO DEVOVLER UN ESTADO DE CODIGO AL MOMENTO DE DAR UNA RESPUESTA SIEMPRE PARA IDENTIFICAR QUE ESTAMOS ENVIANDO, ES UN ESTANDAR
