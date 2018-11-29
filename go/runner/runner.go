package runner

import (
	"fmt"
	"syscall/js"
)

type Runner struct {
	// handleFuncs	map[string]func ([]js.Value, func (string))
	namespace string
}

func NewRunner(namespace string) Runner {
	createNamespace(namespace)

	return Runner{namespace}
}

func createNamespace(namespace string) {
	js.Global().Set(namespace, map[string]interface{}{})
}

type Done func(args ...interface{})

func (r *Runner) HandleFunc(name string, handler func([]js.Value, Done)) {
	js.Global().Get(r.namespace).Set(name, js.NewCallback(func(args []js.Value) {
		fmt.Println("type of last arg is %v", args[len(args)-1].Type())
		if args[len(args)-1].Type() == js.TypeFunction {
			// args[len(args)-1].Invoke()
			fmt.Println("last arg is a function")
		}

		// first arg is name of js callback handler
		// jsCallbackName := args[:1].String()

		// handler(args[1:], func(args ...interface{}) { js.Global().Get(r.namespace).Get(jsCallbackName).Invoke(...args) })
	}))
}
