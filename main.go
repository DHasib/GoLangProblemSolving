package main

import "fmt"

func main() {
	sum := 0
	for v := 1; v < 1_000; v++ { //set range
		if v%3 == 0 || v%5 == 0 {
			sum += v
		}
		
	}
	fmt.Printf("Sum: %d\n", sum)
}