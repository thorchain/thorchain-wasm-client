package helpers

import (
	"fmt"
	"syscall/js"
)

func ParseString(args []js.Value, i int) (string, error) {
	if i < 0 || i >= len(args) {
		return "", fmt.Errorf("Arg %v required, but not present", i)
	}

	if args[i].Type() != js.TypeString {
		return "", fmt.Errorf("Arg %v must be a string, but got type '%v' instead", i, args[i].Type())
	}

	return args[i].String(), nil
}
