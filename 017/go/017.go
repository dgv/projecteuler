// run in http://go-vim.appspot.com/p/djllsHf7tG
package main

import "fmt"
import "strconv"

var numnames = map[int]string{1:"one",
                              2:"two",
                              3:"three",
                              4:"four",
                              5:"five",
                              6:"six",
                              7:"seven",
                              8:"eight",
                              9:"nine",
                              10:"ten",
                              11:"eleven",
                              12:"twelve",
                              13:"thirteen",
                              14:"fourteen",
                              15:"fifteen",
                              16:"sixteen",
                              17:"seventeen",
                              18:"eighteen",
                              19:"nineteen",
                              20:"twenty",
                              30:"thirty",
                              40:"forty",
                              50:"fifty",
                              60:"sixty",
                              70:"seventy",
                              80:"eighty",
                              90:"ninety",
                              100:"hundred",
                              1000:"thousand"}

func spell(n int)int{ 
    if n<21 || (n%10==0 && n<100) {
        return len(numnames[n])
    }	
    s:= strconv.Itoa(n)
    if len(s) == 2{
        dec,_:= strconv.Atoi(string(s[0])+"0")
        uni,_:= strconv.Atoi(string(s[1]))
        return len(numnames[dec])+len(numnames[uni])
    }
    if len(s) == 3{
        hun,_:= strconv.Atoi(string(s[0]))
        t,_:= strconv.Atoi(string(s[1:]))
		if t == 0 {
            return len(numnames[hun])+len("hundred")
		} else {
			return len(numnames[hun])+len("hundredand")+spell(t)
		}
    }
    return len("onethousand")
}
                            
func main() {
    var c int
    for i:=1;i<1001;i++{
    	c += spell(i)
    }
	fmt.Printf("Were used %d letters to count 1-1000",c)
}
