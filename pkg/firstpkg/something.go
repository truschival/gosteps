package firstpkg

import ("fmt")


type Something struct
{
	name string
}

func (f *Something) Printme(){
	fmt.Printf("name: %s \n", f.name)
}

func NewSomething(val string) *Something{
	s := Something{ name: val}
	return &s
}
