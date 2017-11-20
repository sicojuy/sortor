package sortor

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

var cases [][]int

func init() {
	cases = [][]int{
		make([]int, 5),
		make([]int, 10),
		make([]int, 100),
		make([]int, 1000),
		make([]int, 10000),
		make([]int, 1000000),
		//make([]int, 10000000),
	}
}

func initData(arrs []int) {
	rand.Seed(time.Now().UnixNano())
	size := len(arrs)
	for i := 0; i < size; i++ {
		arrs[i] = rand.Intn(math.MaxInt32)
	}
}

func verifyResult(arrs []int) (bool, error) {
	size := len(arrs)
	for i := 1; i < size; i++ {
		if arrs[i-1] > arrs[i] {
			return false, fmt.Errorf("[%d] = %d, [%d] = %d, not increasing", i-1, arrs[i-1], i, arrs[i])
		}
	}
	return true, nil
}

func TestQuickSort(t *testing.T) {
	for _, arrs := range cases {
		initData(arrs)
		QuickSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func TestRadixSort(t *testing.T) {
	for _, arrs := range cases {
		initData(arrs)
		RadixSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func BenchmarkQuickSort100(b *testing.B) {
	b.N = 100
	arrs := make([]int, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		QuickSort(arrs)
	}
}

func BenchmarkRadixSort100(b *testing.B) {
	b.N = 100
	arrs := make([]int, 100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		RadixSort(arrs)
	}
}

func BenchmarkQuickSort10000(b *testing.B) {
	b.N = 10
	arrs := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		QuickSort(arrs)
	}
}

func BenchmarkRadixSort10000(b *testing.B) {
	b.N = 100
	arrs := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		RadixSort(arrs)
	}
}

func BenchmarkQuickSort1000000(b *testing.B) {
	b.N = 1
	arrs := make([]int, 1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		QuickSort(arrs)
	}
}

func BenchmarkRadixSort1000000(b *testing.B) {
	b.N = 3
	arrs := make([]int, 1000000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		RadixSort(arrs)
	}
}
