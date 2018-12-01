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

func ParseObject(args []js.Value, i int) (js.Value, error) {
	if i < 0 || i >= len(args) {
		return js.Value{}, fmt.Errorf("Arg %v required, but not present", i)
	}

	if args[i].Type() != js.TypeObject {
		return js.Value{}, fmt.Errorf("Arg %v must be an object, but got type '%v' instead", i, args[i].Type())
	}

	return args[i], nil
}

func ParseStringProp(arg js.Value, key string, isRequired bool) (string, error) {
	if arg.Get(key).Type() == js.TypeUndefined {
		if isRequired {
			return "", fmt.Errorf("Object must contain (string) prop %v", key)
		}
		return "", nil
	}
	if arg.Get(key).Type() != js.TypeString {
		return "", fmt.Errorf("Object must contain string prop %v, but got type '%v' instead", key, arg.Get(key).Type())
	}
	return arg.Get(key).String(), nil

}
func ParseIntProp(arg js.Value, key string, isRequired bool) (int, error) {
	if arg.Get(key).Type() == js.TypeUndefined {
		if isRequired {
			return 0, fmt.Errorf("Object must contain (int) prop %v", key)
		}
		return 0, nil
	}
	if arg.Get(key).Type() != js.TypeNumber {
		return 0, fmt.Errorf("Object must contain int prop %v, but got type '%v' instead", key, arg.Get(key).Type())
	}
	return arg.Get(key).Int(), nil
}
