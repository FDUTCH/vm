package tokenizer

import "fmt"

type ErrorUnexpected struct {
	Unexpected string
}

func (e ErrorUnexpected) Error() string {
	return fmt.Sprintf("unexpected: %s", e.Unexpected)
}
