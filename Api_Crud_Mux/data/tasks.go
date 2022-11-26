package data

import "crud_GO/models"

// con esto se llena el arreglo de tasks con elementos por determinado para cargar
var TasksData = models.Tasks{
	{
		ID:      1,
		Name:    "Task One",
		Content: "Some Content",
	},
}
