package main

import (
	"flag"
	"fmt"
	"html/template"
	"os"

	"github.com/flosch/pongo2"
)

// the line:
//go:generate go run main.go -arg int -ret int -o map.go
//go:generate go run main.go -arg string -ret int -o mapParallel.go -p true
// works as such: The arg flag represents the function you want to pass in to the map function's argument type
// a.k.a  func Map(arg Type) retType { ... }
var (
	fnArgT   = flag.String("arg", "", "function argument type; ex: int")
	fnRetT   = flag.String("ret", "", "function return type; ex: string")
	parallel = flag.Bool("p", false, "whether or not to generate the parallel version of map;")
	output   = flag.String("o", "map.go", "output file name; ex: mapInt.go")
)

func main() {
	flag.Parse()
	if !*parallel {
		fmt.Println(*fnArgT, *fnRetT, *output)
		val := MapDefine{
			FunctionArgType:    *fnArgT,
			FunctionReturnType: *fnRetT,
		}
		t, _ := template.ParseFiles("map.template")
		f, _ := os.Create(*output)
		t.Execute(f, &val)
	} else {
		fmt.Println(*fnArgT, *fnRetT, *output)
		val := pongo2.Context{
			"FunctionArgType":    *fnArgT,
			"FunctionReturnType": *fnRetT,
		}
		t, _ := pongo2.FromFile("map_parallel.template")
		f, _ := os.Create(*output)
		t.ExecuteWriter(val, f)
	}
	a := make([]int, 10, 100)
	a[0] = 10
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
