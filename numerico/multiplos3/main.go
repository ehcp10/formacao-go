package main

import "fmt"

func main() {
	fmt.Println("Os números divisíveis por 3 são:")
	for i := 1; i <= 100; i++ {
		if (i % 3) == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
