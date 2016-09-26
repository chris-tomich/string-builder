package string_builder

type Word []byte

type StringBuilder struct {
	words []Word
	length int
}

func (sb *StringBuilder) Append(s string) {
	w := Word(s)

	sb.words = append(sb.words, w)
	sb.length = sb.length + len(w)
}

func (sb *StringBuilder) ToString() *string {
	b := make([]byte, 0, sb.length)

	for _, w := range sb.words {
		b = append(b, w)
	}

	return &string(b)
}
