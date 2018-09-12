package main

import "fmt"

var o = "AREA MENGHITUNG"
var types = 10

type jagoan int

var j jagoan

func main() {
	x := 50
	y := 50
	z := ("50 ditambah 50 jadi sangling melengkapi")
	g := x + y
	fmt.Println(o)
	fmt.Println(g)
	fmt.Println(z)
	//jenis apa itu o?
	fmt.Printf("%T\n", o)
	foo()
}

func foo() {
	fmt.Printf("%#x\t%b\t%x\n", types, types, types)

	w := fmt.Sprintf("%#x\t%b\t%x\n", types, types, types)
	fmt.Println(w)
	j = 12
	fmt.Println(j)
	fmt.Printf("%T\n", j)
}
