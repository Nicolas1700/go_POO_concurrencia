package main

import "fmt"

type Person struct {
	name 	string
	age 	int
}

type Employe struct {
	id 		int
}

type FullTimeEmploye struct {
	Person
	Employe
	endDate int
}

// getMessage -> FullTimeEmploye
func (ftEmploye FullTimeEmploye) getMessage() string {
	return "Full Time Employee"
}

type TemporatyEmployee struct {
	Person
	Employe
	taxRate int
}

// getMessage -> TemporatyEmployee
func (tEmploye TemporatyEmployee) getMessage() string {
	return "Temporty Time Employee"
}

// Cambia a interfaz
/*func GetMessage (p Person) {
	fmt.Printf("%s con edad %d\n", p.name, p.age)
}*/

//Interfaz
type PrintInfo interface {
	getMessage() string
}

// Llamar el metodo getMessage -> el cual redirige |
// | al struct conveniente, o fomra "implicita"
func getMessage(p PrintInfo){
	fmt.Println(p.getMessage())
}

func main() {

	ftEmploye := FullTimeEmploye{}
	ftEmploye.name = "Nicolas"
	ftEmploye.age = 18
	ftEmploye.id = 1
	//fmt.Println(ftEmploye)
	//GetMessage(ftEmploye)
	tEmploye := TemporatyEmployee{}
	getMessage(tEmploye)
	getMessage(ftEmploye)

}	
