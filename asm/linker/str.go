package linker

import "fmt"

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
	t.value = val[1 : len(val)-1]
	return nil
}

func (t *StringType) Read(p []byte) (n int, err error) {
	return copy(p, t.value), nil
}
