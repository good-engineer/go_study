/*
this is a code of practicing and learning go language
*/
package main

import (
	"fmt"
	"math"

	"rsc.io/quote"

	_ "time" // use _ to ignore for compile error
)

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())

	/*
	 */
	i := 10
	if i >= 5 {
		fmt.Println("ok")
	}
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	var num int = 1
	if num == 1 {
		fmt.Println("It is!")
	}

	var name string = "Mary"
	var age int = 25

	fmt.Println(name, age)

	var school = "SNU"
	var year int
	year = 2020

	fmt.Println(school, year)

	location := "Seoul"
	fmt.Println("location is", location)

	var country, birthyear = "Iran", 1995

	fmt.Println(country, birthyear)

	var (
		id, phoneNo     = 1, "0107"
		hieght      int = 167
	)
	fmt.Println(id, phoneNo, hieght)
	var x int
	_ = x // to prevent compile error
	var m = -1.34
	fmt.Println(math.Abs(m))
	fmt.Println(math.MaxFloat32)

}
