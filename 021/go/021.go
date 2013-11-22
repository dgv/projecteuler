// run in http://go-vim.appspot.com/p/zTGep34VAB 
package main

import "fmt"

func sum(numbers []int) int {
	s := 0
	for i := 0; i < len(numbers); i++ {
		s += numbers[i]
	}
	return s
}

func divs(n int) []int {
	numbers := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

func main() {
	nums := []int{}
	for i := 1; i < 10000; i++ {
		s := sum(divs(i))
		d := sum(divs(s))
		if d == i {
			if i != s {
				nums = append(nums, d)
			}
		}
	}
	fmt.Printf("The sum of all the amicable numbers under 10000 is %d\n", sum(nums))
}
