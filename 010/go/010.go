// run in http://go-vim.appspot.com/p/Ocrgbr3cg-
package main

import "fmt"

func main() {
	primes := []int{2, 3, 5, 7}
	sum := 17
	n := 7
	p := 7
	for p < 2000000 {
		n++
		for i := 0; i < len(primes); i++ {
			if (primes[i] * primes[i]) > n {
				primes = append(primes, n)
				sum = sum + n
				p = n
				break
			}
			if (n % primes[i]) == 0 {
				break
			}
		}
	}
	fmt.Printf("The sum of all the primes below two million is %d", (sum - n))
}
