// run in http://go-vim.appspot.com/p/ahCSHNQOjI
package main

import "fmt"

func main() {
	// 0..1..1..2..
	// ^        ^
	i := 3
	first := 1
	second := 2
	term := 0
	sum := 2
	for term < 4000000 {
		if i%3 == 0 {
			sum = sum + term
		}
		term = first + second
		first = second
		second = term
		i++
	}
	fmt.Printf("The sum of the even-valued terms is %d", sum)
}
