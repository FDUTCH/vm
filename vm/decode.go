package vm

func Decode(word uint32) (byte, byte, byte, byte, uint16) {
	op := byte(word >> 24)
	dst := byte((word >> 20) & 0x0F)
	src1 := byte((word >> 16) & 0x0F)
	src2 := byte((word >> 12) & 0x0F)
	data := uint16(word & 0xFFF)
	return op, dst, src1, src2, data
}
