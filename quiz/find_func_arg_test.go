package quiz

import (
	"testing"
	"time"
)

func TestFindFuncArgument(t *testing.T) {
	tests := []struct {
		f func(int, int) int
		z int
	}{
		{
			func(x int, y int) int {
				return x + y
			},
			5000,
		},
		{
			func(x int, y int) int {
				return x*x + y
			},
			5000,
		},
	}
	for idx, c := range tests {
		start := time.Now()
		_ = FindFuncArgumentNaive(c.f, c.z)
		t.Logf("native case: #%d \telapsed time: %s", idx+1, time.Now().Sub(start))
		start2 := time.Now()
		_ = FindFuncArgumentBinary(c.f, c.z)
		t.Logf("binary case: #%d \telapsed time: %s", idx+1, time.Now().Sub(start2))
	}
}

func BenchmarkFindFuncArgument(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindFuncArgumentNaive(func(x int, y int) int {
			return x + y
		},
			5000)
	}
}
