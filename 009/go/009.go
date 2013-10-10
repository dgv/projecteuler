// run in http://go-vim.appspot.com/p/L4SALRA_09
package main

import "fmt"

func main() {
	for a := 1; a <= 1000; a++ {
		for b := a; b <= 1000; b++ {
			c := 1000 - a - b
			if c*c == a*a+b*b {
				fmt.Printf("The product of abc is %d", a*b*c)
				break
			}
		}
	}
}
