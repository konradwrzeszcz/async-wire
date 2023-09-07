package main

import "sync"

type Provider[T any] struct {
	wg    sync.WaitGroup
	value T
}

func NewProvider[T any](init func() T) *Provider[T] {
	p := Provider[T]{}
	p.wg.Add(1)

	go func() {
		p.value = init()
		p.wg.Done()
	}()

	return &p
}

func (p *Provider[T]) Value() T {
	p.wg.Wait()

	return p.value
}
