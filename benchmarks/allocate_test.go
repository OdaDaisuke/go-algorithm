package test

import (
	"fmt"
	"testing"
)

func BenchmarkAppend_AllocateEveryTime(b *testing.B) {
	base := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base = append(base, fmt.Sprintf("No%d", i))
	}
}

func BenchmarkAppend_AllocateOnce(b *testing.B) {
	base := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		base[i] = fmt.Sprintf("No%d", i)
	}
}
