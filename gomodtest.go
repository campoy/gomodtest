package gomodtest

import (
	"context"
	"log"
)

// Foo fooes things.
func Foo(ctx context.Context) {
	log.Println("Foo was called")
}

// Bar bars things.
func Bar(ctx context.Context) {
	log.Println("Bar was called")
}
