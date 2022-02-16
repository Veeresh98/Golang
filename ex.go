package main

import "fmt"

func main() {
	avg := func (a int) int {
		return a
	}(12)

	fmt.Println(avg)
			// func parameter (receiver) identifier {code} syntax 

	sum := func (a int , b int) int  {
		return a + b
	}(5,7)

	fmt.Println(sum)

	hey := func (ss string) string  {
		return ss
		
	}("uchiha madara")
	fmt.Println(hey)

	i := bar()
	fmt.Println(i)

	g := i()

	fmt.Println(g)
}

func bar() func () int {
	return func () int {
		return 007
		
	}
	
}

