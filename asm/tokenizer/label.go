package tokenizer

type Label struct {
	name string
}

func (l Label) Name() string {
	return l.name
}
