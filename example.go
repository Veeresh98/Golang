package main

import "fmt"

func fact(nums... int) int {

	tot := 1
	for _,num := range nums{
		tot = tot * num
		
	}
	return tot

}

func main(){
	fmt.Println(fact(1,2,3,4,5,6))

	defer foo()
	boo()
}

func foo(){
	fmt.Println("ffoo")
}

func boo(){
	fmt.Println("boo")
	
	ff()
}

type person struct{
	first string 
	last string
}

type action struct{
	person
	non bool
}

func ff(){

	act := action{
		person :  person{
		
			"naruto",
			"uchiha",
		},
		non : true,
	}
	fmt.Println(act)

	gg()
} 


type shape interface{
	area() float32
	circum() float32

}
type triangle struct{
	length float32	
}

type circle struct{
	radius float32
}

func (s triangle) area() float32{

	return s.length * s.length
	
}

func (c circle) circum() float32{
	return c.radius * c.radius
}

func measure(f shape){
	fmt.Println(f)
	fmt.Println(f.area())
	fmt.Println(f.circum())
}
func gg(){
	s := triangle{length: 10.25}
	c := circle{radius: 5.6}
	

	measure(s)
	measure(c)
}

