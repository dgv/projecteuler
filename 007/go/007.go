// run in http://go-vim.appspot.com/p/85ZKkklUl3
package main

import "fmt"

func main() {
	var primes = []int{2, 3, 5, 7, 11, 13}
	n := 14
	for len(primes) <= 10001 {
		for i := 0; i < len(primes); i++ {
			if (primes[i] * primes[i]) > n {
				primes = append(primes, n)
				break
			}
			if (n % primes[i]) == 0 {
				break
			}
		}
		n++
	}
	fmt.Printf("The 10 001st prime number is %d\n", (n - 1))
}
