// run in http://go-vim.appspot.com/p/AQjpV61dK1
package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func torial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, torial(n.Sub(x, n)))
}

func main() {
	s := 0
	n := torial(big.NewInt(100)).String()
	for i := 0; i < len(n); i++ {
		a, _ := strconv.Atoi(string(n[i]))
		s += a
	}
	fmt.Printf("The sum of the digits in the number 100! is %d", s)
}
