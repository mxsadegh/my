//project eular - problem 3 - Largest prime factor
package main

import (
	"strconv"
	"fmt"
)

func is_reverce(i int) bool{
	str := strconv.Itoa(i)

	var rev_str string
	for _ ,s := range str{
		rev_str = string(s) + rev_str
	}
	i2 , err := strconv.Atoi(rev_str)
	if err != nil {
		fmt.Println("error occure")
	}
	if i2 == i {
	return true
	}else {
		return false
	}

}

var big  =0

func main() {

	for i:=999 ; i>=1;i--{
		for i2:=999;i2>=1 ;i2--  {
			if is_reverce(i2*i) && i2*i > big{
				big = i2*i
			}
		}
	}
	fmt.Println(big)

}
