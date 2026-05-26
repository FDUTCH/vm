package memory

type Flags uint8

const (
	FlagWrite = 1 << iota
	FlagExec
	FlagStack
)

func (r Flags) Write() bool {
	return r&FlagWrite != 0
}

func (r Flags) Exec() bool {
	return r&FlagExec != 0
}

func (r Flags) Stack() bool {
	return r&FlagStack != 0
}
