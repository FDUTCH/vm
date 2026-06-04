package vm

import "math"

var _ Instruction = Equal

// Equal saves performs substraction and saves result to dst.
func Equal(_ *VM, dst *uint64, src1, src2 uint64, _ uint16, _ uint32) {
	*dst = src1 - src2
}

var _ Instruction = Greater

// Greater stores 1 to dst if src1 > src2.
func Greater(_ *VM, dst *uint64, src1 uint64, src2 uint64, flags uint16, _ uint32) {
	switch {
	case (flags & FlagUnsigned) != 0:
		if src1 > src2 {
			*dst = 1
		}
	case (flags & FlagFloat) != 0:
		if math.Float64frombits(src1) > math.Float64frombits(src2) {
			*dst = 1
		}
	default:
		if int64(src1) > int64(src2) {
			*dst = 1
		}
	}
}
