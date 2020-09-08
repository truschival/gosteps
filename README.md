# My first steps in go

Go is amazingly well documented! I remember learning Java and the first contact 
with javadoc (and the sun java libray documentation) it just felt great.
Go takes this to the 2020s and makes it feel really good

To get a glance there is [tour.golang.org](https://tour.golang.org).
Another great resource is [gobyexample](https://gobyexample.com/).

**This is my learning repo - do not take it as reference**

## how to organize a project

Of course this is documented as well! A good start is 
[How to Write Go Code](https://golang.org/doc/code.html)

There is also a complete project layout standard with documentation at
[golang-standards/project-layout](https://github.com/golang-standards/project-layout)

The recommended minimal structure is to have three folders: `pkg` `cmd` and
`test`. 
-  `pkg` contains new packages, so you would have a sub-directory `pkg/mypkg` 
-  `cmd` contains the 'main' package and entry point (possibly with
   sub-directories)
-  `test` contains tests - you can structure them as you wish

An important file is the module description `go.mod` in the root folder. A
module is a set of packages delivered and versioned together.

You can create one with:
```
	go mod init github.com/truschival/mymodule
```


### A Note on Testing

Cool - go has testing built in using the package `testing` and the command `go
test` I put the tests in the separate directory test/* - I don't like to mix
library code with test code. This has some implications.
1. for `go test` to work the test files have to be called `<somename>_test.go`
2. Test files that are executed must be in package main
3. since the tests are in a separate directory they must import the package to
   test i.e. 
   ```
	   package main
	   import (
		"testing"
		"github.com/truschival/gosteps/pkg/firstpkg"
	   )
   ```
   
4. To get coverage for your code (in package firstpkg) you have to explicitly
   tell go test the module path i.e. ```go test
   -coverpkg=github.com/truschival/gosteps/pkg/firstpkg -v
   github.com/truschival/gosteps/test/firsttest```

5. You can actually compile the tests using `go test -c`

## Objects and functions

I have been taught object oriented programming to be the thing to do in my early
days. I even have written tons of C code mimicking OOP techniques. Go is not
object oriented and has a few interesting semantics that feel weird at the first
glance.  But it makes total sense - it is just different. Go uses interfaces and
composition. 

Interfaces are types that combine a set of methods or functions.

You define Interfaces. Just a bunch of functions with a given name and signature
in a custom type interface with curly braces.

You declare a function that consume given interface. I.e. takes an argument of
type interface you just declared.

Then you declare your custom types as struct with their respective properties.
there is no explicit declaration that you have any link to the interface you
declared.  No Java like `class Foo implements Bar` or C++ `class Foo : public
Bar`

To implement the interface you create functions that take the receiver-object of
your type either by value or by pointer and have the same name and signature as
in the interface.

If you implemented all functions for your type that are demanded by the
interface you are done. If you call the function consuming your interface and
your type has not implemented all functions it will not compile.

The weird thing is anything can implicitly satisfy your interface. Even if it
has nothing to do with your initially designed architecture! An object form
another package can "accidentally" implement your interface and you can use it
as such.  The only requirement is that the functions have to be exported
i.e. start the function name with capital letter. This is side effect rarely
discussed in tutorials.  If the functions are not capitalized you will get
strange errors like:

```
   cannot use bychance (type *firstpkg.Something) as type MyInterface in 
   argument to useItf:
   *firstpkg.Something does not implement MyInterface (missing printMe method)
                have firstpkg.printMe()
                want printMe()
```

**Note to myself:** Interface functions should be **C**apitalized

### Some of the stuff I have to get used to:

**You can declare functions for which the receiver struct is passed by value.**
This means this "member function" does create a copy of the struct and does not
manipulate the original. Weird. You have to declare the function with the
receiver as pointer which then modifies the members of the object.  The cool
part here is that you can call either implementation on an value or a pointer to
the object. Go will convert it.

**There is no overloading in Go! you cannot have to functions with the same
name for one type (or free standing)**

**Capitalize function names** if you want them to be exported
