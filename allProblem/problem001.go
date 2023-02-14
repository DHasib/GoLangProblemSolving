package problems

import "fmt"

//Find the sum of all the multiples of 3 or 5 below 1000
func  SummationOfMultiplesValues() {
	sum := 0
	for v := 1; v < 1_000; v++ { //set range
		if v%3 == 0 || v%5 == 0 { //making condition to cheak 3 & 5 multiples
			sum += v // store & addition 3 & 5 multiples
		}
		
	}
	fmt.Printf("Summation Of of 3|5 multiples: %d\n", sum) //show 3 & 5 multiples result the sum
}
