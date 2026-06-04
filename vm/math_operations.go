package vm

import "math"

var _ Instruction = Inc

func Inc(_ *VM, dst *uint64, _ uint64, _ uint64, _ uint16, _ uint32) {
	*dst++
}

var _ Instruction = Dec

func Dec(_ *VM, dst *uint64, _ uint64, _ uint64, _ uint16, _ uint32) {
	*dst--
}

var _ Instruction = AddImmediate

// AddImmediate adds data to dst.
func AddImmediate(_ *VM, dst *uint64, _, _ uint64, _ uint16, data uint32) {
	*dst += uint64(data)
}

var _ Instruction = SubImmediate

// SubImmediate subtracts data from dst.
func SubImmediate(_ *VM, dst *uint64, _, _ uint64, _ uint16, data uint32) {
	*dst -= uint64(data)
}

var _ Instruction = Add

func Add(_ *VM, dst *uint64, src1, src2 uint64, flags uint16, _ uint32) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 + src2
	case (flags & FlagFloat) != 0:
		*dst = math.Float64bits(math.Float64frombits(src1) + math.Float64frombits(src2))
	default:
		*dst = uint64(int64(src1) + int64(src2))
	}
}

var _ Instruction = Sub

func Sub(_ *VM, dst *uint64, src1, src2 uint64, flags uint16, _ uint32) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 - src2
	case (flags & FlagFloat) != 0:
		*dst = math.Float64bits(math.Float64frombits(src1) - math.Float64frombits(src2))
	default:
		*dst = uint64(int64(src1) - int64(src2))
	}
}

var _ Instruction = Mul

// Mul performs multiplication.
func Mul(_ *VM, dst *uint64, src1, src2 uint64, flags uint16, _ uint32) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 * src2
	case (flags & FlagFloat) != 0:
		*dst = math.Float64bits(math.Float64frombits(src1) * math.Float64frombits(src2))
	default:
		*dst = uint64(int64(src1) * int64(src2))
	}
}

var _ Instruction = Div

// Div performs division.
func Div(_ *VM, dst *uint64, src1, src2 uint64, flags uint16, _ uint32) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 / src2
	case (flags & FlagFloat) != 0:
		*dst = math.Float64bits(math.Float64frombits(src1) / math.Float64frombits(src2))
	default:
		*dst = uint64(int64(src1) / int64(src2))
	}
}
