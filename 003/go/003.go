// run in http://go-vim.appspot.com/p/wN3Cy4lDy9
package main

import "fmt"

func main() {
	n := 600851475143
	i := 2
	for i*i < n {
		for n%i == 0 {
			n = n / i
		}
		i++
	}
	fmt.Printf("The largest prime factor of the number 600851475143 is %d", n)
}
