package main

import (
	"fmt"
	"sync"
)

var c = func() (string, error) {
	return "hey!", nil
}
var t = func(a string) {
	fmt.Println(a)
}
var e = func(e error) {
	fmt.Println(e.Error())
}

func main() {
	NewPromise(c).Then(t, e).Value()
}

type Promise struct {
	sync.WaitGroup
	res  string
	err  error
	done bool
}

func (p *Promise) Value() string {
	p.Wait()
	return p.res
}

func NewPromise(f func() (string, error)) *Promise {
	fmt.Println("New Promise...")
	p := &Promise{}
	p.Add(1)
	go func() {
		p.res, p.err = f()
		p.Done()
	}()
	return p
}

func (p *Promise) Then(r func(string), e func(error)) *Promise {
	fmt.Println("Then...")
	go func() {
		p.Wait()
		if p.err != nil {
			e(p.err)
			return
		}
		r(p.res)
	}()
	return p
}
