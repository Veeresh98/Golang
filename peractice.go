package main

import "fmt"

type shape interface{
	area() float32
	circum() float32
}

type square struct{
	length float32
}

type circle struct{
	radius float32
}

func(s square) area() float32{
	return s.length * s.length
}

func(c circle) circum() float32{
	return c.radius * c.radius
}

func measure(g shape) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.circum())
}

func main(){
	s := square{length : 10.54}
	c := circle{radius: 10.21}

	measure(s)
	measure(c)
}