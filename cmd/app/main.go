package main

import (
	"os"
	"plugin"
)

func main() {
	modulePath := "./plugin1.so"
	if _, err := os.Stat(modulePath); os.IsNotExist(err) {
		panic(err)
	}

	p, err := plugin.Open(modulePath)
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
