package main

import "fmt"

func sum(values ... int) int {
	total := 0
	for _, num := range values {
		total += num
	}
	return total
}

func printNombres (nombres ... string){
	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
}

func numExonentes (numero int) (doble, triple, quad int, ){
	doble = numero * 2
	triple = numero * 3
	quad = numero *4
	return
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1,4,5,6))
	printNombres("Nicolas","pedro","Lucas")

	// Uso de retorno con nombre
	fmt.Println(numExonentes(2))

}
