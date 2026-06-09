package linker

import (
	"encoding/binary"
	"fmt"
	"math"
	"strconv"
	"unsafe"
	"vm/asm/tokenizer"
)

type FloatType[T tokenizer.Float] struct {
	value T
}

func (f *FloatType[T]) Name() string {
	return fmt.Sprintf("f%d", unsafe.Sizeof(f.value)*8)
}

func (f *FloatType[T]) Parse(val string) error {
	v, err := strconv.ParseFloat(val, int(unsafe.Sizeof(f.value)*8))
	f.value = T(v)
	return err
}

func (f *FloatType[T]) Value() (uint32, bool) {
	return math.Float32bits(float32(f.value)), unsafe.Sizeof(f.value) == 4
}

func (f *FloatType[T]) Read(p []byte) (n int, err error) {
	if unsafe.Sizeof(0) == 4 {
		binary.LittleEndian.PutUint32(p, math.Float32bits(float32(f.value)))
		return 4, nil
	}
	binary.LittleEndian.PutUint64(p, math.Float64bits(float64(f.value)))
	return 8, nil
}
