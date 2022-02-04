# Manejo de excepciones: DEFER, PANIC & RECOVER

 En GO las excepciones no existen. Tenemos es la función **panic**, que solo debemos usarla cuando queramos que el programa se pare y no deba continuar. 
 ```
 func main() {
    panic("lanzamos función lanic")

    fmt.Println("Este mensaje no se debería mostrar")
}
 ```
 Por otro lado, tenemos la función **defer**, la cual sirve para ejecutar al final del programa la instrucción que le pongamos. Es muy usada para tareas de limpiezas como por ejemplo cerrar conexiones a la base de datos.
 ```
 func main() {
    defer fmt.Println("Esto se mostrará antes de finalizar el programa, justo después del texto ¡Hola mundo!")
    fmt.Println("¡Hola mundo!")
}
 ```
 El resultado de la ejecución anterior sería:
```
¡Hola mundo!
Esto se mostrará antes de finalizar el programa, justo después del texto ¡Hola mundo!
```

La función **defer** unida con la función recover se suele usar para capturar un panic.
```
func main() {
    defer func() {
        r := recover()
        fmt.Println("último panic recuperado:", r)
        fmt.Println("Función defer")
    }()
    fmt.Println("¡Hola mundo!")
    panic("Lanzamos un panic")
}
```
El resultado sería:
```
¡Hola mundo!
último panic recuperado: Lanzamos un panic
Función defer
```
