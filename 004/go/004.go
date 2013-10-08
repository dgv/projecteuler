// run in http://go-vim.appspot.com/p/AS7zIq_fgM
package main

import "fmt"
import "strconv"

func main() {
    d := 0
	for a := 100; a < 1000; a++ {
		for b := a; b < 1000; b++ {
            n := a * b
            c := 0
            palindrome_found := false
			nstr := strconv.Itoa(n)
			for i := 0; i < (len(nstr) / 2); i++ {
				if nstr[0+i] == nstr[len(nstr)-(1+i)] {
					c++
					if c == (len(nstr) / 2) {
                        palindrome_found = true
					}
				}
			}
			if palindrome_found && (n > d) {
				d = n
			}
		}
	}
	fmt.Printf("The largest palindrome is %d", d)
}
