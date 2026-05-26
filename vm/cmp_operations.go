package vm

var _ Instruction = Equal

// Equal saves performs substraction and saves result to dst.
func Equal(_ *VM, dst *uint32, src1, src2 uint32, _ uint16) {
	*dst = src1 - src2
}

var _ Instruction = Greater

// Greater stores 1 to dst if src1 > src2.
func Greater(_ *VM, dst *uint32, src1 uint32, src2 uint32, flags uint16) {
	switch {
	case (flags & FlagUnsigned) != 0:
		if src1 > src2 {
			*dst = 1
		}
	case (flags & FlagFloat) != 0:
		if *float(&src1) > *float(&src2) {
			*dst = 1
		}
	default:
		if *signed(&src1) > *signed(&src2) {
			*dst = 1
		}
	}
}
