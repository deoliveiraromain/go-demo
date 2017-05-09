package main

import (
	"fmt"
	"time"
)

func compute(n int, c chan int) {
	fmt.Println("Calcul n° ", n)
	time.Sleep(1 * time.Second)
	c <- n
}
func main() {
	c := make(chan int)
	for i := 1; i <= 10; i++ {
		go compute(i, c)
	}
	for i := 1; i <= 10; i++ {
		fmt.Println("End of calcul n°", <-c)
	}
}
