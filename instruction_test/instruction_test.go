package instruction_test

import (
	"testing"
	"vm/asm"
	"vm/vm"
)

func TestEncodeAndDecode(t *testing.T) {
	target := uint64(0x0123456789abcdef)
	op, dst, src1, src2, flags, data := vm.Decode(target)
	t.Logf("op: %x, dst: %x, src1: %x, src2: %x, flags: %x, data: %x", op, dst, src1, src2, flags, data)
	result := asm.Encode(op, dst, src1, src2, flags, data)
	if result != target {
		t.Fatalf("results does not match (target: %x, result: %x)", target, result)
	}
}
