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
	fnArgT = flag.String("arg", "", "function argument type; ex: int")
	fnRetT = flag.String("ret", "", "function return type; ex: string")
	output = flag.String("o", "map.go", "output file name; ex: mapInt.go")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Flags:\n")
	flag.PrintDefaults()
}

func main() {
	//flag.Usage = Usage
	flag.Parse()
	fmt.Println(*fnArgT, *fnRetT, *output)

	val := MapDefine{
		FunctionArgType:    *fnArgT,
		FunctionReturnType: *fnRetT,
	}

	t, _ := template.ParseFiles("map.template")
	f, _ := os.Create(*output)
	t.Execute(f, &val)
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
func MapUnsafe(args []interface{}, fn func(interface{}) interface{}) []interface{} {
	output := make([]interface{}, len(args))
	for i, v := range args {
		output[i] = fn(v.(interface{}))
	}
	return output
}
