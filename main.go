package main

import (
	"context"
	"fmt"
)

type IFoo interface {
	Bar()
}

type Foo struct {
}

func (f Foo) Baz(ctx context.Context) {
	fmt.Println("hello world")
}
