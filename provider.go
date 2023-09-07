package main

import (
	"github.com/stephennancekivell/go-future/future"
)

type Provider[T any] struct {
	f future.Future[T]
}

func NewProvider[T any](init func() T) Provider[T] {
	f := future.New[T](init)
	return Provider[T]{f: f}
}

func (p Provider[T]) Value() T {
	return p.f.Get()
}
