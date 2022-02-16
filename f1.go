package main

import "fmt"

func main() {
	fmt.Println("hello world")

	foo()
}

func foo() {
	fmt.Println("hello my name is veereshgo run ")

	b := "luffy"

	fmt.Println("the value of a is :", "\t", b)

	boo()
}

var (
	fname string
	sname string
	age   int
	area  float32
)

func boo() {

	fname = "monkey"
	fname = "naruto"
	sname = "dragon"
	age = 43
	area = 54.45

	fmt.Println(fname, sname, age, area)

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	num := 400

	if num%4 == 0 {
		fmt.Println("year is a leap year")
	} else {
		fmt.Println("year is not a leap year ")

	}

	type chichi string
	var fry chichi

	fry = "i love fried chicken very well "
	fmt.Println(fry)
	fmt.Printf("%T", fry)
	fmt.Println("\n")

	goo := []int{2, 23, 5, 6, 7, 8, 5, 4, 67, 7}
	for inn, va := range goo {
		fmt.Println(inn, va)
	}

	for j := 65; j <= 125; j++ {
		fmt.Printf("%#U", j)
	}

	day := "tue"

	switch {

	case day == "mon":
		fmt.Println("monday")

	case day == "tue":
		fmt.Println("tuesday")
		fallthrough
	case day == "wed":
		fmt.Println("i watch anime")
	}

	m := map[string]string{
		"fruit":       "apple",
		"yello fruit": "mangoes",
	}
	fmt.Println(m)

	base()
}

type person struct {
	f1  string
	mn  string
	ln  string
	age int
}

func base() {

	p1 := person{
		f1:  "monkey",
		mn:  "d",
		ln:  "dragon",
		age: 25,
	}
	fmt.Println(p1)
	fmt.Println(p1.ln)

	fo()
}

func fo() {
	p2 := struct {
		fname string
		age   int
	}{
		fname: "naruto",
		age:   20,
	}
	fmt.Println(p2)

	for i:= 0; i <6;i++{
		if i %2  ==0 {
		fmt.Println("even")
	}

	}	

	xi := []int{1,2,3,4,5,6}
	nums := goo(xi...)
	fmt.Println(nums)
	
} 

func goo(nums... int){
	total:=1
	for _,num := range nums {
		total = total * num
	}
	fmt.Println(total)
}
