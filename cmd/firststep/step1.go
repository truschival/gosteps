package main

import (
	"fmt"
	"github.com/truschival/gosteps/pkg/firstpkg"
)

func main() {
	fmt.Println("step1 go!")
	var name string = "Thomas"
	fmt.Println(firstpkg.Banner(name))
	longbanner,err := firstpkg.BannerRepeat(3, &name) 
	if err {
		fmt.Println(longbanner)
	}
}
