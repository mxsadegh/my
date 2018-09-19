// project euler - problem 10
//اشتباهم این بود که محدودیت int رو در نظر نگرفتم باید int64 استفاده می کردم

package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("please wait")
	var tedad int64=5
	var ok bool
	for i:=int64(3);;i=i+2{
		for j:=int64(math.Sqrt(float64(i)))+1;j>=3;j-- {
			if i%j==0 {
				ok = false
				break
			}else {
				ok = true
			}

		}

		if i>=int64(2000000){

			break
		}
		if ok{
			tedad += i
		}

	}
 fmt.Println(tedad)
}