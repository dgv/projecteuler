// run in http://go-vim.appspot.com/p/1z60-UBpsF
package main

import "fmt"

var (
	prime_list = []int{2, 3, 5, 7, 11, 13, 17, 19, 23}
	lastn      = 23
)

func isprime(n int) bool {
	isprime := false
	for i := 0; i < len(prime_list); i++ {
		if prime_list[i]*prime_list[i] > n {
			break
		}
		if (n % prime_list[i]) != 0 {
			isprime = true
			break
		}
	}
	return isprime
}

func prime(x int) int {
	for len(prime_list) <= x {
		lastn = lastn + 1
		if isprime(lastn) {
			prime_list = append(prime_list, lastn)
		}
	}
	return prime_list[x]
}

func num_factors(n int) int {
	div := 1
	x := 0
	for n > 1 {
		var c = 1
		for (n % prime(x)) == 0 {
			c = c + 1
			n = n / prime(x)
		}
		x = x + 1
		div = div * c
	}
	return div
}

func main() {
	n := 0
	for i := 1; i <= 1000000000; i++ {
		n = i * (i + 1) / 2
		if num_factors(n) > 500 {
			fmt.Printf("The value of the first triangle number to have over five hundred divisors is %d", n)
			break
		}
	}
}
