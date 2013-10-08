// run in http://go-vim.appspot.com/p/cmBb7yKW_T
package main

import "fmt"

func main() {
	n := 2520
	for n > 1 {
		if n%1 == 0 &&
			n%2 == 0 &&
			n%3 == 0 &&
			n%4 == 0 &&
			n%5 == 0 &&
			n%6 == 0 &&
			n%7 == 0 &&
			n%8 == 0 &&
			n%9 == 0 &&
			n%10 == 0 &&
			n%11 == 0 &&
			n%12 == 0 &&
			n%13 == 0 &&
			n%14 == 0 &&
			n%15 == 0 &&
			n%16 == 0 &&
			n%17 == 0 &&
			n%18 == 0 &&
			n%19 == 0 &&
			n%20 == 0 {
			break
		}
		n++
	}
	fmt.Printf("The smallest positive number that is evenly divisible by all of the numbers from 1 to 20 is %d", n)
}
