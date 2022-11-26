package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"crud_GO/data"
	"crud_GO/models"

	"github.com/gorilla/mux"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task

	//con esto manejamos las entradas y salidas
	reqBody, err := io.ReadAll(r.Body) //ESTO DEVUELVE UN ERROR O LOS DATOS

	if err != nil {
		fmt.Fprintf(w, "Insert a Valid Task Data")
	}

	//Desde request body se lo voya  asignar a newTask
	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(data.TasksData) + 1

	data.TasksData = append(data.TasksData, newTask)

	//esto especifica que devovlemos un JSON
	w.Header().Set("Content-Type", "application/json") //esto deberia estar desde el inicio // ver otro codigo
	w.WriteHeader(http.StatusCreated)                  //develve el estado
	json.NewEncoder(w).Encode(newTask)

}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data.TasksData)
}

func GetOneTask(w http.ResponseWriter, r *http.Request) {
	//vars es un arreglo
	vars := mux.Vars(r)                     //extrae las variables de esa peticion, lo gaurdamos ene l vars que contiene varias variables
	taskID, err := strconv.Atoi(vars["id"]) //strconv, con esto formateamos strings que tenemos, el atoi pasa el string a entero

	if err != nil {
		return
	}

	//recorrer una lista se hace de esta forma, el _ funciona como id y el task como la variable que devuelve\
	//si quisieramos utiliar el indice cambiar _ por i
	for _, task := range data.TasksData {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(task)
		}
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])
	var updatedTask models.Task

	if err != nil {
		fmt.Fprintf(w, "Invalid ID")
	}

	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Please Enter Valid Data")
	}
	json.Unmarshal(reqBody, &updatedTask)

	for i, t := range data.TasksData {
		if t.ID == taskID {

			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			updatedTask.ID = t.ID
			data.TasksData = append(data.TasksData, updatedTask)

			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updatedTask)
			fmt.Fprintf(w, "The task with ID %v has been updated successfully", taskID)
		}
	}

}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "Invalid User ID")
		return
	}

	for i, t := range data.TasksData {
		if t.ID == taskID {
			data.TasksData = append(data.TasksData[:i], data.TasksData[i+1:]...)
			fmt.Fprintf(w, "The task with ID %v has been remove successfully", taskID)
		}
	}
}
