package linker

import (
	"fmt"
	"strings"
)

type StringType struct {
	value string
}

func (t *StringType) Value() (uint32, bool) {
	return 0, false
}

func (t *StringType) Name() string {
	return "str"
}

func (t *StringType) Parse(val string) error {
	if len(val) < 2 {
		return fmt.Errorf("invalid value for str: %s", val)
	}
	t.value = r.Replace(val[1 : len(val)-1])
	return nil
}

func (t *StringType) Read(p []byte) (n int, err error) {
	return copy(p, t.value), nil
}

var r = strings.NewReplacer(
	`\n`, "\n",
	`\t`, "\t",
	`\r`, "\r",
	`\0`, "\x00",
	`\x00`, "\x00",
	`\\`, "\\",
)
