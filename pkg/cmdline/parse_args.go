package cmdline

import (
	"fmt"
)

const foo string = "foo"
const bar string = "bar"
const baz string = "baz"

func ParseArgs(arglist []string) bool {
	object := arglist[0]
	switch object {
	case foo:
		fmt.Printf("> %s\n", foo)
		return true

	case bar:
		fmt.Printf("> %s\n", bar)
		return true

	case "baz":
		fmt.Printf("> %s\n", baz)
		return true

	default:
		return false
	}
}
