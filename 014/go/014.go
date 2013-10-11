// run in http://go-vim.appspot.com/p/yaJ8831s31
package main

import "fmt"

var a = []int{}

func seq(n int) int {
	a = append(a, n)
	if n > 1 {
		if n%2 > 0 {
			n = 3*n + 1
			return seq(n)
		} else {
			n = n / 2
			return seq(n)
		}
	}
	return n
}

func main() {
	var i, n, c int
	for i < 1000000 {
		a = []int{}
		seq(i)
		if len(a) > c {
			c = len(a)
			n = i
		}
		i++
	}
	fmt.Printf("The number which starting number, under one million, \nthat produces the longest chain is %d", n)
}
