package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello workld")

	var a float32

	a = 10.56
	fmt.Println(a)

	var b int = 56
	fmt.Println(b)

	c := "Hello all"
	fmt.Println(c)

	var d = "hey"
	fmt.Println(int(a),d) 

	boo()
	

}
var(
	fn string
	ln string
	age int
)

var ss,gg ="hello","jhon"
func boo(){
	fn = "luffy"
	ln = "sanji"
	age = 21
	fmt.Println(fn,ln,age)
	fmt.Println((ss) + (gg))
	

	foo()


}
func foo(){

	h := "hey"
	j := 10
	fmt.Println(h +  " " + fmt.Sprint(j))


	fmt.Println("enter your name")


	num := 20

	if  num%2 ==0 {
		fmt.Println("it's divided")

	}else{
		fmt.Println("not dividing")
	}

	str :="Monkey D Dragon"

	for i := range str{
		fmt.Println(i)
	}

	const(
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1,c2,c3)

	for k:= 65;k<=97;k++{
		fmt.Printf("%#U",k)
	}

	day := "tue"

	switch{
	case day =="mon":
		fmt.Println("it's monday")
		fallthrough

	case day =="tue":
		fmt.Println("tuesday")
		fallthrough
		

	case day =="wed":
		fmt.Println("wednesday")
		
	}

	lis := []int{1,2,3,4,5}
	fmt.Println(lis)
	lis = append(lis, 6,7,8)
	fmt.Println(lis[2])

	gl := []int{}
	fmt.Println(gl)


	m := map[string]int{
		"Naruto" : 17,
		"Kira" : 20,
		"L" : 35,
	}
	fmt.Println(m["L"])

	ee(56,67,"jhon")
}
func ee(s,v int,ss string){
	fmt.Println("hello",s + v,ss)

	 i := 6
	 switch{
	 case i == 1:
		if i < 6{
			fmt.Println("print the numbers",i)
		}else{
			fmt.Println("none")
		}
		fmt.Println("yes this it's one. ")
		fallthrough

	 case i == 5:
		if i < 10{
			fmt.Println("yes print ",i)
		}else{
			fmt.Println("nothing",i)
		}
		fmt.Println("print nothing")
		fallthrough

	case i == 6:
		if i ==6{
			fmt.Println("let's print this")
		}else{
			fmt.Println("let's print something")
		}

	 }
	 s1 := wo("james")
	 fmt.Println(s1)
	 

	
	
	 
}
func wo(st string)string{
	return fmt.Sprintln("hello",st)

	ggg()


}

type human struct{
	f string
	s string
}

type animal struct{
	breed1 string
	breed2 string 
	breed3 string 

}

func ggg(){
	per1 := human{
		f : "Naruto",
		s : "Big mom",
	}
	fmt.Println(per1)

	an := animal{
		breed1 : "lion",
		breed2 : "tiger",
		breed3 : "wolf",
		
	}
	fmt.Println(an)

	h()
}

type react struct{
	fs string
	ls string
	ag int
}

func (ra react) int{
	return "hello",ra.fs + ra.ls, "and my age is ",ra.age

}
func h(){
	ra := react{"mokey","luffy",12}
}
	





