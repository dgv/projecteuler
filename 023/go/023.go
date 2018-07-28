// run in https://go-vim.appspot.com/p/lP3-rUPKUa
package main

func sumDivs(n int) int {
	sum, p := 1, 2
	for p*p <= n && n > 1 {
		if n%p == 0 {
			j := p * p
			n /= p
			for n%p == 0 {
				j *= p
				n /= p
			}
			sum *= j - 1
			sum /= p - 1
		}
		if p == 2 {
			p = 3
		} else {
			p += 2
		}
	}
	if n > 1 {
		sum *= n + 1
	}
	return sum
}

func main() {
	limit := 28124
	abundants := make([]int, 6965)
	ai := 0
	for i := 2; i < limit; i++ {
		if i < sumDivs(i) - i {
			abundants[ai] = i
			ai++
		}
	}
	sumAb := make([]bool, limit)
	for _, v := range abundants {
		for _, w := range abundants {
			if sum := v + w; sum < limit {
				sumAb[sum] = true
			}
		}
	}
	sum := 0
	for i := range sumAb {
		if !sumAb[i] {
			sum += i
		}
	}
	println("The sum of all the positive integers which cannot be written as the sum of two abundant numbers is",sum)
}
