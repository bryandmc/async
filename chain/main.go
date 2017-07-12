package main

import "fmt"

var c = func(i interface{}) interface{} {
	fmt.Println(i)
	return i.(int) + 10
}

func main() {
	funs := []func(interface{}) interface{}{c, c, c}
	out := ChainUnsafe(funs, 10)
	fmt.Println(out)
}

func ChainUnsafe(fns []func(interface{}) interface{}, firstArg interface{}) interface{} {
	var lastOut = fns[0](firstArg)
	for c, f := range fns[1:] {
		fmt.Println(c)
		lastOut = f(lastOut)
	}
	return lastOut
}
