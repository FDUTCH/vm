package vm

var _ Instruction = Inc

func Inc(_ *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	*dst++
}

var _ Instruction = Dec

func Dec(_ *VM, dst *uint32, _ uint32, _ uint32, _ uint16) {
	*dst--
}

var _ Instruction = AddUint12Immediate

// AddUint12Immediate adds data to dst.
func AddUint12Immediate(_ *VM, dst *uint32, _, _ uint32, data uint16) {
	*dst += uint32(data)
}

var _ Instruction = SubUint12Immediate

// SubUint12Immediate subtracts data from dst.
func SubUint12Immediate(_ *VM, dst *uint32, _, _ uint32, data uint16) {
	*dst -= uint32(data)
}

var _ Instruction = Add

func Add(_ *VM, dst *uint32, src1, src2 uint32, flags uint16) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 + src2
	case (flags & FlagFloat) != 0:
		*float(dst) = *float(&src1) + *float(&src2)
	default:
		*signed(dst) = *signed(&src1) + *signed(&src2)
	}
}

var _ Instruction = Sub

func Sub(_ *VM, dst *uint32, src1, src2 uint32, flags uint16) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 - src2
	case (flags & FlagFloat) != 0:
		*float(dst) = *float(&src1) - *float(&src2)
	default:
		*signed(dst) = *signed(&src1) - *signed(&src2)
	}
}

var _ Instruction = Mul

// Mul performs multiplication.
func Mul(_ *VM, dst *uint32, src1, src2 uint32, flags uint16) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 * src2
	case (flags & FlagFloat) != 0:
		*float(dst) = *float(&src1) * *float(&src2)
	default:
		*signed(dst) = *signed(&src1) * *signed(&src2)
	}
}

var _ Instruction = Div

// Div performs division.
func Div(_ *VM, dst *uint32, src1, src2 uint32, flags uint16) {
	switch {
	case (flags & FlagUnsigned) != 0:
		*dst = src1 / src2
	case (flags & FlagFloat) != 0:
		*float(dst) = *float(&src1) / *float(&src2)
	default:
		*signed(dst) = *signed(&src1) / *signed(&src2)
	}
}

var signed = cast[uint32, int32]
var float = cast[uint32, float32]
