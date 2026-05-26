package main

import (
	"vm/asm"
	"vm/memory"
	"vm/vm"
)

func main() {
	RunVm()
}

func RunVm() {
	data := `
		ldi r1 0 // null терминатор
		ldi r2 0 // ссылка на текущий символ
		subi r3 4 // уменьшаем регистер на который будем прыгать.
		ldb r4 r2 S // loading char from static region
		inc r2 // incrementing pointer
		equal r5 r4 r1 // comparing chars
		jumpc r3 r5 // jumping if r5 != 0
		write r6 r2 S // пишим 
		ldi r1 1 // загружаем 1 в r1
		wstatus r1 // пишем r1 как статус
	`

	some, err := asm.Compile(data)
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
		Data: []byte("Hello World!!!\u0000"),
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
