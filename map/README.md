A simple 'map' implementation.

Usage:

edit main.go to the comment //go:generate ... to satisfy the types needed for your map function or use the MapUnsafe() function. This uses interface{} as all the types and therefore cannot do type checking. It is strongly recommended that you generate a map function for your types as to get the benefits of type safety.

