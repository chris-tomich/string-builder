package string_builder

import (
	"testing"
	"fmt"
)

func TestStringBuilder(t *testing.T) {
	sb := StringBuilder{}
	sb.Append("This")
	sb.Append(" ")
	sb.Append("is")
	sb.Append(" ")
	sb.Append("a")
	sb.Append(" ")
	sb.Append("test.")

	newString := sb.ToString()

	fmt.Println(*newString)
}

func AddSpace(s string) string {
	return s + " "
}

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := AddSpace("This")
		s = s + AddSpace("is")
		s = s + AddSpace("a")
		s = s + AddSpace("test.")

		fmt.Sprintln(s)
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sb := StringBuilder{}
		sb.Append("This")
		sb.Append(" ")
		sb.Append("is")
		sb.Append(" ")
		sb.Append("a")
		sb.Append(" ")
		sb.Append("test.")

		fmt.Sprintln(*(sb.ToString()))
	}
}
