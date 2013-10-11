// run in http://go-vim.appspot.com/p/7VIeIXCbVT
// http://mathworld.wolfram.com/BinomialCoefficient.html
package main

import "fmt"
import "math/big"

func torial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, torial(n.Sub(x, n)))
}

func main() {
	//(a+b)!/a!(a+b-a)! [a=20,b=20]
	x := big.NewInt(1)
	y := big.NewInt(1)
	y.Mul(torial(big.NewInt(20)), torial(big.NewInt(20)))
	x.Div(torial(big.NewInt(40)), y)
	fmt.Println(x)
}
