package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"
)

// the line:
//go:generate go run main.go -arg int -ret int -o map.go
// works as such: The arg flag represents the function you want to pass in to the map function's argument type
// a.k.a  func Map(arg Type) retType { ... }
var (
	fnArgT   = flag.String("arg", "", "function argument type; ex: int")
	fnRetT   = flag.String("ret", "", "function return type; ex: string")
	parallel = flag.Bool("p", false, "whether or not to generate the parallel version of map;")
	output   = flag.String("o", "map.go", "output file name; ex: mapInt.go")
)

var c = func(i interface{}) interface{} {
	return i.(int) + 10
}

func main() {
	flag.Parse()
	fmt.Println(*fnArgT, *fnRetT, *output)

	val := MapDefine{
		FunctionArgType:    *fnArgT,
		FunctionReturnType: *fnRetT,
	}

	t, _ := template.ParseFiles("map.template")
	f, _ := os.Create(*output)
	t.Execute(f, &val)
	// m := MapUnsafeParallel(c, 1, 2, 3, 4)
	// q := MapUnsafe(c, 1, 2)
	// fmt.Println(m)
	// fmt.Println(q)
}

// MapDefine is used when generating type-safe implementations of 'Map' for use
// with any types. This is a safer alternative to MapUnsafe which does no type checking
// for you at compile time.
type MapDefine struct {
	FunctionArgType    string
	FunctionReturnType string
}

// MapUnsafe works like a normal functional programming high order 'Map' function.
// It takes a list of input, applies the function to each item in the list
// and returns another list with the results.
func MapUnsafe(fn func(interface{}) interface{}, args ...interface{}) []interface{} {
	output := make([]interface{}, len(args))
	for i, v := range args {
		output[i] = fn(v.(interface{}))
	}
	return output
}

// MapUnsafeParallel is the same as MapUnsafe but it executes all iterations concurrently.
// Very useful for complex/time consuming functions.
func MapUnsafeParallel(fn func(interface{}) interface{}, args ...interface{}) []interface{} {
	output := make([]interface{}, len(args))
	collector := make(chan interface{}, len(args))
	for _, v := range args {
		go func(i interface{}) {
			collector <- fn(i)
			return
		}(v)
	}
	for j := 0; j < len(args); j++ {
		output[j] = <-collector
	}
	return output
}
