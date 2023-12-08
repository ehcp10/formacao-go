package main

import "fmt"

func main() {
	fmt.Println("Os números divisíveis por 3 são:")
	for i := 1; i <= 100; i++ {
		if (i%3) == 0 && (i%5) == 0 {
			fmt.Printf("PinPan ")
		} else if (i % 3) == 0 {
			fmt.Printf("Pin ")
		} else if (i % 5) == 0 {
			fmt.Printf("Pan ")
		} else {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Printf("\n")
}
