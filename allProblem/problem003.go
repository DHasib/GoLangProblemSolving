package problems

import ( 
	"fmt"
	"math"
)
//What is the largest prime factor of the number 600851475143?


//This function used for define Largest primeFactors
func findlargest(factors []int) int {
    max := factors[0]
    for _, factor := range factors {
        if factor > max {
            max = factor
        }
    }
    return max
}
//This function used for define prime factors
func findPrimeFactors(n int) []int {
    factors := make([]int, 0)
    for i := 2; i <= int(math.Sqrt(float64(n))); i++ {//find the square root[Sqrt] of the number and convert into integer number
        if n%i == 0 {
            factors = append(factors, i)
            n /= i
            i--
        }
    }
    if n > 0 {
        factors = append(factors, n)
    }
    return factors
}

//This function used for input Numbers to calculate largest Prime Factors
func InputNumbersToFindLargestPrimenumber(n int) {
    // n := 600851475143
    largestPrimeFactor := findlargest(findPrimeFactors(n))
    fmt.Printf("largest prime factor of the number %d-\n [%d]\n", n, largestPrimeFactor)
}