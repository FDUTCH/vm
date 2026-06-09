package linker

import (
	"encoding/binary"
	"fmt"
	"strconv"
	"unsafe"
	"vm/asm/tokenizer"
)

type IntegerType[T tokenizer.Integer] struct {
	Unsigned bool
	value    T
}

func (i *IntegerType[T]) Name() string {
	ch := 'i'
	if i.Unsigned {
		ch = 'u'
	}
	return fmt.Sprintf("%c%d", ch, unsafe.Sizeof(i.value)*8)
}

func (i *IntegerType[T]) Parse(val string) error {
	var err error
	size := int(unsafe.Sizeof(i.value) * 8)
	if i.Unsigned {
		var v uint64
		v, err = strconv.ParseUint(val, 0, size)
		i.value = T(v)
		return err
	}
	var v int64
	v, err = strconv.ParseInt(val, 0, size)
	i.value = T(v)
	return err

}

func (i *IntegerType[T]) Value() (uint32, bool) {
	return uint32(i.value), unsafe.Sizeof(i.value) <= 4
}

func (i *IntegerType[T]) Read(p []byte) (n int, err error) {
	size := int(unsafe.Sizeof(i.value))
	switch size {
	case 1:
		p[0] = byte(i.value)
	case 2:
		binary.LittleEndian.PutUint16(p, uint16(i.value))
	case 4:
		binary.LittleEndian.PutUint32(p, uint32(i.value))
	case 8:
		binary.LittleEndian.PutUint64(p, uint64(i.value))
	}
	return size, nil
}
