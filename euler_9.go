package main

import (
	"fmt"
	//"strings"
	)

func main() {
	a,b,c :=4,5,6
	for ;a<500;a++{
		b=5
		for ;b<500;b++{
			c=6
			for ; c<500;c++{
				if a*a+b*b==c*c{
					fmt.Printf("\r a=%d b=%d c=%d",a,b,c)
					if a+b+c==1000{
						goto here
					}
				}
			}
		}
	}
here:
	fmt.Println("\n",a*b*c)
}