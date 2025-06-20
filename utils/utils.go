package utils

import (
	"strings"

	s "golang.org/x/sync/singleflight"
)

var Group s.Group

type Singleflight[T any] struct {
	Group *s.Group
	Key   string
}

func GenKey(parts ...string) string {
	return strings.Join(parts, "::")
}

func (single *Singleflight[T]) ProcessWrapper(fn func() (T, error)) (T, error) {
	wrapperFn := func() (interface{}, error) {
		return fn()
	}

	res, err, _ := single.Group.Do(single.Key, wrapperFn)

	if err != nil {
		var zero T
		return zero, err
	}

	return res.(T), nil
}

func (single *Singleflight[T]) Forget(keys ...string) {
	for _, key := range keys {
		single.Group.Forget(key)
	}
}
