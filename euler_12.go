//project euler - problem 12
// خیلی خواستم ابتکار به خرج بدم یه چیز عجیب به جای
// جزر n نوشتم n/i-1+1
//واسه همین سرعت برنامه کم بود که اینجوری بیشتر شد
//همین بد حساب کردن جزر باعث شد که مقسوم هارو نصف در نظر بگیرم به جز 1 و n که دوتا حساب شدن
package main

import "fmt"

func triangle(n int64) int64{
	num := int64(0)
	for i:=int64(1) ; i<=n ; i++ {

		num += i
	}
	return num
}

func maghsoom (n int64) int64{
	total :=int64(1)
	for i:=int64(2) ; i<math.Sqrt(float64(n)) ; i++ {

		if n%i == 0 {
			total +=1
		}
	}

	return (total)*2
}

func main() {

	for i:=int64(2);;i++{

		if maghsoom(triangle(i)) > 500 {
			fmt.Println(triangle(i))

			break
		}
	}

}