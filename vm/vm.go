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

	program []uint64

	status uint64

	stack   memory.Region
	static  memory.Region
	dynamic memory.Region

	registry *Registry
	opcodes  [256]Instruction

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

	if len(program.Data)%8 != 0 {
		return nil, fmt.Errorf("invalid program alignment")
	}

	p := make([]uint64, 0, len(program.Data)/8)
	for i := 0; i < len(program.Data); i += 8 {
		p = append(p, binary.LittleEndian.Uint64(program.Data[i:i+8]))
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

	// setting instruction set if it is nil.
	if vm.registry == nil {
		OptionInstructionSet(DefaultRegistry())(vm)
	}
	return vm, nil
}

func (vm *VM) Tick() {
	// fetching instruction.
	word := vm.program[vm.registers[memory.ProgramCounterReg]]

	// decoding instruction.
	op, d, s1, s2, flags, data := Decode(word)

	// setting source & destination registers.
	dst := &vm.registers[d]
	src1 := vm.registers[s1]
	src2 := vm.registers[s2]

	// executing instruction.
	vm.opcodes[op](vm, dst, src1, src2, flags, data)

	// incrementing program counter.
	vm.registers[memory.ProgramCounterReg]++
}

func (vm *VM) Status() *uint64 {
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
