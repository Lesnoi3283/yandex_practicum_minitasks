package main

import (
	"bytes"
	"strings"
	"testing"
)

func BenchmarkFibo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FiboRecursive(20)
	}
}

func BenchmarkAlignRight(b *testing.B) {
	s := "some string"
	length := 300
	lead := rune('_')

	b.Run("ARPlus", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlignRightPlus(s, length, lead)
		}
	})

	b.Run("ARBuf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlignRightBuf(s, length, lead)
		}
	})

	b.Run("ARStrRepeat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			AlignRightStrRepeat(s, length, lead)
		}
	})

}

func AlignRightPlus(s string, length int, lead rune) string {
	for len(s) < length {
		s = string(lead) + s
	}
	return s
}

func AlignRightBuf(s string, length int, lead rune) string {
	buf := bytes.Buffer{}
	for i := 0; i < length-len(s); i++ {
		buf.WriteRune(lead)
	}
	buf.WriteString(s)
	return buf.String()
}

func AlignRightStrRepeat(s string, length int, lead rune) string {
	if len(s) < length {
		return strings.Repeat(string(lead), length-len(s)) + s
	}
	return s
}
