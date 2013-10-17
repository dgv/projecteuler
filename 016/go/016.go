// run in http://go-vim.appspot.com/p/OyUf8pvpjW
package main

import "fmt"
import "math/big"
import "strconv"

func main() {
	n := big.NewInt(1).Exp(big.NewInt(2), big.NewInt(1000), big.NewInt(0)).String()
	r := 0
	for i := 0; i < len(n); i++ {
		a, _ := strconv.Atoi(string(n[i]))
		r += a
	}
	fmt.Println("The sum of the digits of the number 2^1000 is", r)
}
