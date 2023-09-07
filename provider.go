package main

import (
	"github.com/stephennancekivell/go-future/future"
)

type Provider[T any] struct {
	F future.Future[T]
}

func NewProvider[T any](init func() T) Provider[T] {
	f := future.New[T](init)
	return Provider[T]{F: f}
}
