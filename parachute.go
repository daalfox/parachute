package parachute

import (
	"errors"

	"golang.org/x/sync/singleflight"
)

type Result[T any] struct {
	Val    T
	Err    error
	Shared bool
}

type Parachute[T any] struct {
	sf singleflight.Group
}

func (g *Parachute[T]) Do(key string, fn func() (T, error)) (T, error, bool) {
	val, err, shared := g.sf.Do(key, func() (any, error) {
		return fn()
	})

	v, ok := val.(T)
	if !ok {
		var zero T
		return zero, errors.New("unexpected type in singleflight result"), shared
	}

	return v, err, shared
}

func (g *Parachute[T]) DoChan(key string, fn func() (T, error)) <-chan Result[T] {
	sfChan := g.sf.DoChan(key, func() (any, error) {
		return fn()
	})

	ch := make(chan Result[T], 1)
	v := <-sfChan
	val, ok := v.Val.(T)
	if !ok {
		return nil
	}

	ch <- Result[T]{
		Val:    val,
		Err:    nil,
		Shared: false,
	}

	return ch
}

func (g *Parachute[T]) Forget(key string) {
	g.sf.Forget(key)
}
