//project eular - problem 5 - Smallest multiple
package main

import "fmt"
// با دادن مقدار 2520 به شمارشگر i
// و شروع j از 11 خیلی سرعت برنامه بیشتر شد
func main()  {
	i:=2520
	for j :=11;j<=20 ; j++{
		if i%j==0 {
			continue
	}else {
		i += 2520
		j=11
		continue
		}

	}
	fmt.Println(i)
}