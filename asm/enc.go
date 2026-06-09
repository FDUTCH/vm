package asm

import "vm/asm/tokenizer"

func Encode(op byte, dst byte, src1 byte, src2 byte, flags uint16, data uint32) uint64 {
	word := uint64(0)
	word |= uint64(op) << 56
	word |= uint64(dst&0xf) << 52
	word |= uint64(src1&0xf) << 48
	word |= uint64(src2&0xf) << 44
	word |= uint64(flags&0xfff) << 32
	word |= uint64(data)
	return word
}

func EncodeInstructionToken(i tokenizer.Instruction) uint64 {
	return Encode(i.ID, i.Dst, i.Src1, i.Src2, i.Flags, i.Data.Value())
}
