package instruction

import (
	"testing"
	"vm/asm"
	"vm/memory"
	"vm/vm"
)

func TestRunVM(t *testing.T) {
	data := `
		SubImmediate r3 4 // уменьшаем регистер на который будем прыгать.
		LoadByte r4 r2 $S // loading char from static region
		Inc r2 // incrementing pointer
		Equal r5 r4 r1 // comparing chars
		JumpRelConditional r3 r5 // jumping if r5 != 0
		Write r6 r2 $S // пишим 
		LoadImmediate r1 1 // загружаем 1 в r1
		WriteStatus r1 // пишем r1 как статус
	`
	reg := vm.DefaultRegistry()

	some, err := asm.Compile(data, reg)
	if err != nil {
		panic(err)
	}

	program := memory.Region{
		Name:  "program",
		Flags: memory.FlagExec,
		Data:  some,
	}
	static := memory.Region{
		Name: "static",
		Data: []byte("Hello World!!!\n\u0000"),
	}
	stack := memory.Region{
		Name:  "stack",
		Flags: memory.FlagStack,
		Data:  make([]byte, 0, 1024),
	}

	machine, err := vm.NewVM([]memory.Region{program, static, stack})
	if err != nil {
		panic(err)
	}
	status := machine.Status()

	for *status == 0 {
		machine.Tick()
	}
}
