// run in http://go-vim.appspot.com/p/3EtmNkiC1i
package main

import "fmt"

func main() {
	y := 100
	d := 1 // week:0-7 sun:0
	c := 0
	for y > 0 {
		//jan
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//feb
		if y%4 == 0 { // last
			if d = (29 + d) % 7; d == 0 {
				c++
			}
		} else {
			if d == 0 {
				c++
			}
		}
		//mar
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//apr
		if d = (30 + d) % 7; d == 0 {
			c++
		}
		//may
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//jun
		if d = (30 + d) % 7; d == 0 {
			c++
		}
		//jul
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//aug
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//set
		if d = (30 + d) % 7; d == 0 {
			c++
		}
		//out
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		//nov
		if d = (30 + d) % 7; d == 0 {
			c++
		}
		//dec
		if d = (31 + d) % 7; d == 0 {
			c++
		}
		y--
	}

	fmt.Printf("There's %d sundays at first of month, between 1 Jan 1901 and 31 Dec 2000", c)
}
