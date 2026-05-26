package vm

const (
	FlagStaticRegion = 1 << iota // S - flag
	FlagNoneRegion               // N - flag
	FlagStackRegion              // K - flag
	FlagUnsigned                 // U - flag
	FlagFloat                    // F - flag
	FlagInvert                   // I - flag
)
