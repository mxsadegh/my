//project eular - problem 3 - Largest prime factor
package main

import (
	"fmt"
)

func main() {
//برای مقایسه اعداد باید واحدشون یکی باشه
//وهر نوع عددی میتونه توی حلقه for و if بیاد
// به شرطی که با هم هم واحد باشن

	n:=int64(600851475143)
	var amel int64
	for i:=int64(2) ;  i <= n ;{
		if n%i ==0 {
			amel = i
			n /= i
			continue

		}else {
			i++
		}
	}
	fmt.Println(amel)

}
