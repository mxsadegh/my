package main

import (
	"fmt"
)

func hesab(n int) int {
	var chain = 0
	ok := false
	for ok == false {
		if n == 1 {
			ok = true
		} else if n%2 == 0 {
			n /= 2
			chain++
		} else {
			n = n*3 + 1
			chain++
		}
	}
	return chain
}

func main() {
	var biggest = 1
	for i := 1; i < 1000000; i++ {
		if hesab(i) > biggest {
			biggest = hesab(i)
			fmt.Println(i)
		}
	}

}
