package vm

type Instruction = func(vm *VM, dst *uint32, src1, src2 uint32, data uint16)

var opcodes [256]Instruction

func init() {
	opcodes[OpInc] = Inc
	opcodes[OpDec] = Dec
	opcodes[OpLoadUInt16Immediate] = LoadUInt16Immediate
	opcodes[OpLoadUint32] = LoadUint32
	opcodes[OpLoadUint16] = LoadUint16
	opcodes[OpLoadByte] = LoadByte
	opcodes[OpMove] = Move
	opcodes[OpMoveToAddrUint32] = MoveToAddrUint32
	opcodes[OpMoveToAddrUint16] = MoveToAddrUint16
	opcodes[OpMoveToAddrByte] = MoveToAddrByte
	opcodes[OpAddUint12Immediate] = AddUint12Immediate
	opcodes[OpSubUint12Immediate] = SubUint12Immediate
	opcodes[OpAdd] = Add
	opcodes[OpSub] = Sub
	opcodes[OpMul] = Mul
	opcodes[OpDiv] = Div
	opcodes[OpJump] = Jump
	opcodes[OpJumpConditional] = JumpConditional
	opcodes[OpEqual] = Equal
	opcodes[OpGreater] = Greater
	opcodes[OpCall] = Call
	opcodes[OpRet] = Ret
	opcodes[OpPushImmediate] = PushImmediate
	opcodes[OpPushUint16] = PushUint16
	opcodes[OpPopUint16] = PopUint16
	opcodes[OpPushUint32] = PushUint32
	opcodes[OpPopUint32] = PopUint32
	opcodes[OpPush] = Push
	opcodes[OpPop] = Pop
	opcodes[OpWrite] = Write
	opcodes[OpRead] = Read
	opcodes[OpWriteStatus] = WriteStatus
	opcodes[OpLoadStatus] = LoadStatus
}

const (
	OpInc = iota
	OpDec
	OpLoadUInt16Immediate
	OpLoadUint32
	OpLoadUint16
	OpLoadByte
	OpMove
	OpMoveToAddrUint32
	OpMoveToAddrUint16
	OpMoveToAddrByte
	OpAddUint12Immediate
	OpSubUint12Immediate
	OpAdd
	OpSub
	OpMul
	OpDiv
	OpJump
	OpJumpConditional
	OpEqual
	OpGreater
	OpCall
	OpRet
	OpPushImmediate
	OpPushUint16
	OpPopUint16
	OpPushUint32
	OpPopUint32
	OpPush
	OpPop
	OpWrite
	OpRead
	OpWriteStatus
	OpLoadStatus
)
