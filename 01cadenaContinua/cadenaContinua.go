package main

import "fmt"

/*
Este
*/
func main() {
	/*Problema no entendido:
	Al momento de asignar una variable Scanln en la misma
	línea detecta un error ya que Scanln solo libera dos
	variables una de error y otra de tipo int.
	*/
	var palabra string
	var num int8
	fmt.Println("Indica una palabra")
	fmt.Scanln(&palabra)
	fmt.Println("Indica un número")
	fmt.Scanln(&num)

	count := 0
	for i := 0; i < int(num); i++ {
		count = count % len(palabra)
		fmt.Print(string(palabra[count]))
		count++
	}
}
