package assignments

const (
	funcDef = `package p1

import "fmt"

func Greet() {
	fmt.Println("Hello Src2Crs!")
}
`
)

func ExampleReadAndCallFunction() {
	ReadAndCallFunction(funcDef, "p1.Greet()")

	// Output:
	// Hello Src2Crs!
}
