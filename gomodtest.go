package gomodtest

import (
	"context"
	"log"

	"github.com/campoy/gomodtest/stuff"
)

// Foo fooes things.
func Foo(ctx context.Context) {
	log.Println("Foo was called")
}

// Bar bars things.
func Bar(ctx context.Context) {
	log.Println("Bar was called")
}

func Answer() {
	log.Println("The answer is", stuff.Answer)
}
