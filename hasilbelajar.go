/**
package main

import (
	"fmt"
)

func main() {
x := 42
y := "James Bond"
z := true
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(x, y, z)
}
**/
/**
package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

**/
/**
package main

import (
	"fmt"
)
var x int = 42
var y string = "James Bond"
var z bool = true
func main() {

	
	s := fmt.Sprintf("%v\t%v\t%v\n", x, y, z)
	fmt.Println(s)
}

**/
/**
package main

import (
	"fmt"
)
type boy int
var x boy
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

**/
/**
package main

import (
	"fmt"
)
type boy int
var x boy
var y int
func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

**/
/**
package main

import "fmt"

var x bool
func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
}

**/
/**
package main

import "fmt"

func main() {
	a := 7
	b := 42
	fmt.Println(a == b)
}

**/