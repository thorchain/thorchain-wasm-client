package runner

import (
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

func (r *Runner) HandleFunc(name string, handler func([]js.Value) []interface{}) {
	js.Global().Get(r.namespace).Set(name, js.NewCallback(func(args []js.Value) {
		// check if last arg is a function, if yes invoke it
		if len(args) > 0 && args[len(args)-1].Type() == js.TypeFunction {
			res := handler(args[0 : len(args)-1])
			args[len(args)-1].Invoke(res...)
		} else {
			handler(args)
		}
	}))
}
