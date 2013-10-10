// run in http://go-vim.appspot.com/p/GGOUgvCfAJ
package main

import "fmt"

func main() {
	var primes = []int{2, 3, 5, 7, 11, 13}
	count := len(primes)
	n := primes[5] + 1
	for count < 10001 {
		remainders := 0
		for i := 0; i < len(primes); i++ {
			if (n % primes[i]) == 0 {
				continue
			} else {
				remainders++
			}
			if remainders == len(primes) {
				count++
				primes = append(primes, n)
				break
			}
		}
		n++
	}
	fmt.Printf("The 10 001st prime number is %d\n", (n - 1))
}
