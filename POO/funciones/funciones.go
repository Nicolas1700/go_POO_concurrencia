package main

import (
	"fmt"
	"time"
)

func main() {

	// Fun anonima sobre gorutines
	c := make(chan int)
	go func(){
		fmt.Println("Se inico el proceso...")
		time.Sleep(5 * time.Second)
		fmt.Println("Se ejecuto :)")
		c <- 1
	}()
	<- c
	
}