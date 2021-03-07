package main

import "fmt"

var (
	i2 = 2
	j2 = 5
	s2 = "a & b"
	b2 = true
)
func main()  {
	fmt.Println("variable")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	fmt.Println(i2, j2, s2, b2)
}

func variableZeroValue()  {
	var i int
	var s string
	fmt.Printf("%d %q", i, s)
	fmt.Println()
}

func variableInitialValue()  {
	var i,j int = 1,2
	var s string = "hello"
	fmt.Println(i,j,s)
}

func variableTypeDeduction()  {
	var i, j, s = 1, 2, "hello2"
	fmt.Println(i,j,s)
}

func variableShorter()  {
	i, j, s := 1, 2, "hello3"
	j = 3
	fmt.Println(i,j,s)
}
