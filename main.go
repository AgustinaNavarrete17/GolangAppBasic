package main

import (
	"fmt"
	"net/http"
)

/*Defino un tipo de dato llamado tarea y dentro de un struct defino que datos tendrá mi tarea: */
type task struct {
	/*Como vamos a estar respondiendo como un json al cliente, vamos a decirle que vamos a responderle con los mismos
	nombres que estos datos:
	(Recordar poner los nombres luego de json con comillas) */
	ID      int    `json:"ID" `
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

//Creo un tipo donde estaran listadas todas las tareas, es decir, sera un arreglo:
type allTask []task

//Aqui defino la lista de todas las tareas y la guardo en una variable a la cual nombro "tasks":
var tasks = allTask{
	{ /*Aqui tengo un solo objeto, con el que puedo hacer tareas basicas de nuestra API, como consultar,
		borrar, actualizar el contenido de nuestros datos.
		PERO puedo crear mas objetos con otros datos */
		ID:      1,
		Name:    "Task one",
		Content: "Some content"},
}

/*Esta funcion solo la creo para llamarla en la funcion main, en la variable router
y recibe de parte del HandleFunc dos parametros: un objeto w que sera http.ResponseWriter", que esta relacionado con
la respuesta que envio y despues un objeto r que va a estar relacionado con la peticion del usuario
y sera http.Request*/
func indexRoute(w http.ResponseWriter, r *http.Request) {
	//Y respondere imprimiendo lo siguiente:
	fmt.Fprintf(w, "Bienvenido a mi API")
}

func main() {
	/*Creo un enrutador con la biblioteca MUX:
	  Este enrutador especificara y corregira cual es la ruta url correcta en caso de que se ingrese una incorrecta,
	  y coloque la correcta en el navegador del usuario
	  Y guardo el enrutador en una variable llamada router, que la creo e inicializo: */
	router := mux.newRouter().StrictSlash(true)

	/*uso esa variable para definir una ruta a través de HandleFunc, la cual va a recibir un nombre
	  y que es lo que quiero que haga cuando visiten esa URL:
	*/
	//Entre () le digo: cuando visiten la url ppal (que lo represento con esto: "/"), quiero que ejecute lo siguiente (lo que va luego de la coma)
	router.HandleFunc("/", indexRoute)

}
