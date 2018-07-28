// run in https://go-vim.appspot.com/p/cTXONnMqqV
package main

func main() {
	digs := 1000
	n1 := make([]int, digs, digs)
	n2 := make([]int, digs, digs)
	var a, b []int
	n1[0], n2[0] = 1, 1
	for i := 3; ; i++ {
		if i%2 == 0 {
			a = n1
			b = n2
		} else {
			a = n2
			b = n1
		}
		for n := 0; n < digs; n++ {
			if n > 0 && a[n-1] > 9 {
				a[n] += int(a[n-1] / 10)
				a[n-1] %= 10
			}
			a[n] += b[n]
		}
		if a[digs-1] != 0 {
			println("The index of the first term in the Fibonacci sequence to contain 1000 digits", i)
			break
		}
	}
}
