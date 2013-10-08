// run in http://go-vim.appspot.com/p/CFaGbsEHSS
package main

import "fmt"
import "math"

func main() {
	var a, b float64
	a = 0
	b = 0
	for i := 1; i <= 100; i++ {
		a = a + math.Pow(float64(i), float64(2))
		b = b + float64(i)
	}
	c := math.Pow(b, float64(2)) - a
	fmt.Printf("The difference between the sum of the squares of the first one hundred natural numbers and the square of the %d", int(c))
}
