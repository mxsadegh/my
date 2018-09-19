package main

import "fmt"

func main() {
	var sum ,sqr int
	for i:=1 ; i<101 ; i++{
		sqr += i*i
		sum += i
	}
	fmt.Printf("sum is  : %d and sqr is: %d \n",sum,sqr)
	fmt.Println(sum*sum - sqr)
}