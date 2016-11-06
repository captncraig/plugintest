package main

import (
	"fmt"
	"plugin"
)

func main() {
	fmt.Println("HOST")
	p, err := plugin.Open("plugin.so")
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	fmt.Println(f)
	f.(func())()
}
