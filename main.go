package main

import (
	"fmt"
	_ "strconv"
	_ "time"
)

type Employes struct {
	id int
	name string
	desicion bool
}

func (e *Employes) SetId(id int){
	e.id = id
}
func (e *Employes) SetName(name string){
	e.name = name
}

func (e *Employes) GetId() int {
	return e.id
}
func (e *Employes) GetName() string {
	return e.name
}
// 4
// Crear strcut con funcion, ref val struct y ret &
func NewEmploye (id int, name string, desicion bool) *Employes {
	return &Employes{
		id: 		id,
		name: 		name,
		desicion: 	true,
	}
}

func main() {
	
	// 1
	e := Employes{}
	fmt.Println(e)
	// 2
	e2 := Employes{
		id: 1,
		name: "Nicolas",
		desicion: true,
	}
	fmt.Println(e2)
	// 3
	// De esta forma nos devulve una referencia &
	// Por ello se va imprimir con *valor
	e3 := new(Employes)
	fmt.Println(*e3)
	// 4
	// Creando una funcion con la accion de crear
	// la struct, mirar la funcion
	e4 := NewEmploye(3,"Pedrito",true)
	fmt.Println(*e4)
}

