package model

import (
	"testing"
)

func BenchmarkNewPoem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewPoem("abc", "def")
	}
}
