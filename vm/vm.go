package vm

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"vm/memory"
)

type VM struct {
	registers memory.Registers

	program []uint32

	status uint32

	stack   memory.Region
	static  memory.Region
	dynamic memory.Region

	stdIn  io.Reader
	stdOut io.Writer
}

func NewVM(regions []memory.Region, options ...Option) (*VM, error) {
	var (
		program memory.Region
		static  memory.Region
		dynamic memory.Region
		stack   memory.Region
	)

	for _, rg := range regions {
		switch {
		case rg.Flags.Write():
			dynamic = rg
		case rg.Flags.Exec():
			program = rg
		case rg.Flags.Stack():
			stack = rg
		default:
			static = rg
		}
	}

	if len(program.Data)%4 != 0 {
		return nil, fmt.Errorf("invalid program alignment")
	}

	p := make([]uint32, 0, len(program.Data)/4)
	for i := 0; i < len(program.Data); i += 4 {
		p = append(p, binary.LittleEndian.Uint32(program.Data[i:i+4]))
	}

	vm := &VM{
		program: p,
		stack:   stack,
		static:  static,
		dynamic: dynamic,
		stdIn:   os.Stdin,
		stdOut:  os.Stdout,
	}
	for _, op := range options {
		op(vm)
	}
	return vm, nil
}

func (vm *VM) Tick() {
	// fetching instruction.
	word := vm.program[vm.registers[memory.ProgramCounterReg]]

	// decoding instruction.
	op, d, s1, s2, data := Decode(word)

	// setting source & destination registers.
	dst := &vm.registers[d]
	src1 := vm.registers[s1]
	src2 := vm.registers[s2]

	// executing instruction.
	opcodes[op](vm, dst, src1, src2, data)

	// incrementing program counter.
	vm.registers[memory.ProgramCounterReg]++
}

func (vm *VM) Status() *uint32 {
	return &vm.status
}

func (vm *VM) GetRegion(flags uint16) *memory.Region {
	switch {
	case flags&FlagNoneRegion != 0:
		return nil
	case flags&FlagStaticRegion != 0:
		return &vm.static
	case flags&FlagStackRegion != 0:
		return &vm.stack
	}
	return &vm.dynamic
}
