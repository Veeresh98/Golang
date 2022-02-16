package main

import (
	"fmt"
	"math"
)


type geometry interface {
    area() float64
    perim() float64

}

type rect struct {
    width, height float64

}

type circle struct {
    radius float64

}

func (r rect) area() float64 {
    return r.width * r.height

}

func (r rect) perim() float64 {
    return 2*r.width + 2*r.height

}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius

}


func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())

}

func main() {
    r := rect{width: 3, height: 4}
    c := circle{radius: 5}

    measure(r)
	measure(c)

	
}



type done interface{
	even() int16
	year() int16
	fact() int16
}

type numbers struct{
	values int16
}
type num struct {
	val int16
}

func (n numbers) even() int16 {
	if n.values %2 ==0{
		fmt.Println("even number :",n.values)
	}else{
		fmt.Println("odd numbner :",n.values)
	}
}

func (nn num) year() int16  {
	if nn.val %4 == 0{
		fmt.Println("leap year")
	}else{
		fmt.Println("not a leap year! ")
	}
}

func (){
	
}