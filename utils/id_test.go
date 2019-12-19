package utils

import (
	"testing"
)

func BenchmarkGenerateID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateID()
	}
}

func BenchmarkGenerateShortID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateShortID()
	}
}
