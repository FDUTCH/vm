package tokenizer

func Expect(str string, tok rune, expected ...rune) error {
	for _, val := range expected {
		if tok == val {
			return nil
		}
	}
	return ErrorUnexpected{str}
}
