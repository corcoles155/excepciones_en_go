package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"log"
)


func main()  {
	
	ejemploDefer()

	ejemploPanic()

}	

func ejemploDefer()  {
	fileName := "prueba.txt"
	fichero, err := os.Open(fileName)

	defer fichero.Close() //Antes de salir de la función debe de ejecutar fichero.close()

	if err != nil {
		fmt.Println("error abriendo el fichero")
		os.Exit(1) //finaliza la ejecución del propgrama
	}
}

func ejemploPanic()  {
	//Se ejecutará después del panic. defer SIEMPRE se ejecuta al terminar el programa
	defer func ()  {
		//Creamos una función anónima para poder ejecutar varias instrucciones
		//recover() obtiene el resultado del último panic, si no hubo un panic devuelve nil
		reco := recover()
		if (reco != nil) {
			//log.Fatalf muestra un texto en la consola y escribe en el log
			log.Fatalf("Ocurrió un error que generó Panic \n %v", reco) //el verbo "%v" permite colocar una variable
		}
	}() //Es una función anónima que no devuelve nada
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1") //panic fuerza un error, mostrará "Se encontro el valor 1" y cerrará el programa
	}
}
