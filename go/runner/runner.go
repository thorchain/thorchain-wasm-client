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

func (r *Runner) HandleFunc(name string, handler func([]js.Value) (interface{}, error)) {
	js.Global().Get(r.namespace).Set(name, js.NewCallback(func(args []js.Value) {
		// TODO may recover from panics instead, which is bad style, but otherwise the wasm module will crash if e. g. a string value was not provided from JS and types were not checked properly: https://stackoverflow.com/a/22222070/6694848

		// check if last arg is a callback, if yes invoke it
		if len(args) > 0 && args[len(args)-1].Type() == js.TypeFunction {
			res, err := handler(args[0 : len(args)-1])
			if err != nil {
				args[len(args)-1].Invoke(fmt.Sprintf("%v", err), res)
			} else {
				args[len(args)-1].Invoke(nil, res)
			}
		} else {
			_, err := handler(args)
			if err != nil {
				fmt.Println("Received err but did not define callback:", err)
			}
		}
	}))
}
