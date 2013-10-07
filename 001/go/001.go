// run in http://go-vim.appspot.com/p/z3XKwNvWY9
package main

import "fmt"

func main() {
	multiples := 0
	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			multiples += i
		}
	}
	fmt.Printf("The sum of all the multiples of 3 or 5 below 1000 is %d", multiples)
}
