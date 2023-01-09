package assignments

import (
	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

func ReadAndCallFunction(funcDef, call string) bool {
	i := interp.New(interp.Options{})
	i.Use(stdlib.Symbols)

	_, err := i.Eval(funcDef)
	if err != nil {
		panic(err)
	}

	_, err = i.Eval(call)
	if err != nil {
		panic(err)
	}

	return true
}
