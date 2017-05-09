package main

import (
	"fmt"
	"time"
)

func compute(n int) {
	fmt.Println("Calcul n° ", n)
	time.Sleep(1 * time.Second)
	fmt.Println("End of calcul n°", n)
}
func main() {
	for i := 1; i <= 10; i++ {
		go compute(i)
	}
}
