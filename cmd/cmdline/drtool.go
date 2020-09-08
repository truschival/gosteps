package main

import (
	"flag"
	"fmt"
	"os"
	"github.com/truschival/gosteps/pkg/cmdline"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Use this magic like this:\n")
	flag.PrintDefaults()
}


func main() {
	// Default initialize to env variable
	mojo := flag.String("s", os.Getenv("SALT"), "secret mojo")
	flag.Usage = usage;
	flag.Parse()
	if *mojo == "" {
		panic("No valid mojo argument")
	}

	args := flag.Args()
	fmt.Printf("Args:\t%v \nMojo:\t%s \n", args[:], *mojo)
	
	if !cmdline.ParseArgs(args){
		flag.PrintDefaults()
	}
}

