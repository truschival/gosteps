package main

import (
	"fmt"
	"math"
	"github.com/truschival/gosteps/pkg/firstpkg"
)

// Declare an interface type i.e.
// functions to be provided
type MyInterface interface {
	Printme()
	// getType(format string)
	// does not work - no overload in go!
}

// declare some types
// no interface to see here
type myTypeF struct {
	value float64
}

type myTypeI struct {
	value int64
}

// Go does have constructors per se. it is idiomatic to declare
// Constructor functions
func newMyTypeF(val float64) *myTypeF {
	r := myTypeF{value: val}
	return &r
}

func newMyTypeI(val int64) *myTypeI {
	r := myTypeI{value: val}
	return &r
}

// Functions i.e. methos for these types
// names and signatures according to interface I want to implement
// but apart form function name nothing indicates I use an interface
func (f myTypeF) Printme() {
	fmt.Printf("Float: %E\n",f.value)
}

func (i myTypeI) Printme() {
	fmt.Printf("Int : %d\n",i.value)
}

// Overload by argument does not work in go
// func (f myTypeF) Printme(format string) {
// 	fmt.Printf(format, f.value)
// }

// func (i myTypeI) Printme(format string) {
// 	fmt.Printf(format, i.value)
// }


// I can define addional methods for my types
// Really interesting: methods also get called with object by value!
// i.e. (i myTypeI) creates a copy of i and modifies the value !?!?
func (i *myTypeI) specialFunc(something int) {
	i.value += int64(something)
}

// Some function that requires the interfaces
func useItf(itf MyInterface) {
	itf.Printme()
}

func main() {
	fmt.Printf("test overload:\n")
	i := myTypeI{value: 1e2}
	j := newMyTypeI(44)
	x := new(myTypeI)
	// and the cool stuff: go converts value to pointer
	i.specialFunc(-23)
	f := myTypeF{value: math.Pi}

	// call on pointer to struct
	x.value = 42
	x.specialFunc(100)

	useItf(f)
	useItf(i)
	useItf(j)
	useItf(x)

	// use an object from another pkg that does not know about myInterface
	bychance := firstpkg.NewSomething("widget")
	useItf(bychance) 
}
