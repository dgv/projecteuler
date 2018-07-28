// run in https://go-vim.appspot.com/p/3W_w2OqY6R
package main

func torial(i int) int {
	if i < 2 {
		return 1
	}
	return i * torial(i-1)
}

func main() {
	n := 999999
	m := 9
	l := m + 1
	nums := make([]int, l)
	for i := 0; i < l; i++ {
		nums[i] = i
	}

	var fact, dig, dig_ int
	number := make([]int, l)
	for i := m; i >= 0; i-- {
		fact = torial(i)
		dig_ = int(n / fact)
		dig = nums[dig_]
		number[m-i] = dig
		nums = append(nums[:dig_], nums[dig_+1:]...)
		n -= fact * dig_
	}
	print("The millionth lexicographic permutation of the digits 0-9 is ")
	for _, d := range number {
		print(d)
	}
}
