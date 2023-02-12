package problems

import "fmt"

//By considering the terms in the Fibonacci sequence whose values do not exceed 4 million(4,000,000), find the sum of the even-valued terms.
func SummationOfEvenFibonacciValues(){
	v1 := 0
    v2 := 1
    sum := 0 
    const max = 4000000
    for v2 < max { //conditionally check Fibonacci max value & sequencially generater Fibonacci output
        buff := v1 + v2 
        if v2%2 == 0 {//conditionally check & Sum of all Fibonacci even output 
            sum += v2
        }
        v1 = v2
        v2 = buff
    }
    fmt.Printf("Summation Of Even Fibonacci Values(: [%d]\n", sum)//show result
}
