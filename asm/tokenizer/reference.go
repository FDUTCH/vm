package tokenizer

import "fmt"

type Reference[T Number] struct {
	value    T
	hasValue bool
	name     string
}

func NewReferenceWithValue[T Number](value T) Reference[T] {
	return Reference[T]{value: value, hasValue: true}
}

func NewReferenceWithName[T Number](name string) Reference[T] {
	return Reference[T]{name: name}
}

func (s *Reference[T]) Value() T {
	return s.value
}

func (s *Reference[T]) NeedsValue() bool {
	return !s.hasValue && s.name != ""
}

func (s *Reference[T]) Link(symbol map[string]uint32) error {
	if !s.NeedsValue() {
		return nil
	}

	val, has := symbol[s.name]
	if !has {
		return fmt.Errorf("no symbol called: %s found", s.name)
	}

	s.SetValue(T(val))
	return nil
}

func (s *Reference[T]) SetValue(val T) {
	s.value = val
	s.hasValue = true
}

func (s *Reference[T]) Name() string {
	return s.name
}

type Integer interface {
	~int8 | ~int16 | ~int32 | ~int64 |
		~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	Integer | Float
}
