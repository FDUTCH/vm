package vm

import "io"

type Option = func(vm *VM)

func OptionStdIn(reader io.Reader) Option {
	return func(vm *VM) {
		vm.stdIn = reader
	}
}

func OptionStdOut(writer io.Writer) Option {
	return func(vm *VM) {
		vm.stdOut = writer
	}
}

func OptionInstructionSet(set *Registry) Option {
	return func(vm *VM) {
		vm.registry = set
		vm.opcodes = set.instructions
	}
}
