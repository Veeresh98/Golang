package main

import "fmt"

func main(){
	v := 10
	fmt.Println(v)

	sum := func (a int , b int) int {
		return a + b
	}(12,32)

	fmt.Println(sum)

	 type person struct{
		 name string
		 age int
	 }

	 humans := person{
		 name : "monkey d luffy",
		 age :19,
	 }
	 fmt.Println("his name " , humans.name, " and age is ", humans.age)

	 s := []int{1,2,3,4,5,6}
	 fmt.Println(s)

	 for c,d := range s{
		 fmt.Println(c,d)
	 }
}