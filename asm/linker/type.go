package linker

import (
	"io"
	"vm/asm/tokenizer"
)

type Type interface {
	// Name returns name of the type.
	Name() string
	// Parse parses string representation into value.
	Parse(val string) error
	// Value returns value (represented in 32 bits) and true if it can fit in 32 bits.
	Value() (uint32, bool)
	io.Reader
}

func DefaultTypeRegistry() TypeRegistry {
	r := make(TypeRegistry, 11)
	RegisterIntegerType[int8](r, false)
	RegisterIntegerType[int16](r, false)
	RegisterIntegerType[int32](r, false)
	RegisterIntegerType[int64](r, false)

	RegisterIntegerType[uint8](r, true)
	RegisterIntegerType[uint16](r, true)
	RegisterIntegerType[uint32](r, true)
	RegisterIntegerType[uint64](r, true)

	RegisterFloatType[float32](r)
	RegisterFloatType[float64](r)

	RegisterStringType(r)
	return r
}

func RegisterIntegerType[T tokenizer.Integer](r TypeRegistry, unsigned bool) {
	r[(&IntegerType[T]{Unsigned: unsigned}).Name()] = func() Type {
		return &IntegerType[T]{Unsigned: unsigned}
	}
}

func RegisterFloatType[T tokenizer.Float](r TypeRegistry) {
	r[(&FloatType[T]{}).Name()] = func() Type {
		return &FloatType[T]{}
	}
}

func RegisterStringType(r TypeRegistry) {
	r[(&StringType{}).Name()] = func() Type {
		return &StringType{}
	}
}

type TypeRegistry map[string]func() Type

func (r TypeRegistry) Get(name string) (Type, bool) {
	fn, ok := r[name]
	if !ok {
		return nil, false
	}
	return fn(), true
}
