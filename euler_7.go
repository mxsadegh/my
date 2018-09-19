package main

import "fmt"

func main() {
	fmt.Println("please wait")
	tedad :=2
	var ok bool
	for i:=3;;i=i+2{
		for j:=i-1;j>=3;j-- {
			if i%j==0 {
				ok = false
				break
			}else {
				ok = true
			}

		}
		if ok{
			tedad++
		}
		if tedad==10001{
			fmt.Println("\r",tedad,i)
			break
		}
		fmt.Printf("\r %d %d",tedad,i)
	}

}