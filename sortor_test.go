package sortor

import (
	"fmt"
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
		arrs[i] = rand.Intn(size)
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

func TestInsertionSort(t *testing.T) {
	cases := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 10000),
	}
	for _, arrs := range cases {
		initData(arrs)
		InsertionSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	cases := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 10000),
	}
	for _, arrs := range cases {
		initData(arrs)
		SelectionSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func TestHeapSort(t *testing.T) {
	cases := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 10000),
		make([]int, 1000000),
	}
	for _, arrs := range cases {
		initData(arrs)
		HeapSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func TestQuickSort(t *testing.T) {
	cases := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 10000),
		make([]int, 1000000),
	}
	for _, arrs := range cases {
		initData(arrs)
		QuickSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func TestRadixSort(t *testing.T) {
	cases := [][]int{
		make([]int, 10),
		make([]int, 100),
		make([]int, 10000),
		make([]int, 1000000),
		make([]int, 10000000),
	}
	for _, arrs := range cases {
		initData(arrs)
		RadixSort(arrs)
		if b, err := verifyResult(arrs); !b {
			t.Errorf("array size is %d, sort error: %s", len(arrs), err)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.N = 2
	arrs := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		InsertionSort(arrs)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	b.N = 2
	arrs := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		SelectionSort(arrs)
	}
}

func BenchmarkHeapSort(b *testing.B) {
	b.N = 10
	arrs := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		initData(arrs)
		b.StartTimer()
		HeapSort(arrs)
	}
}

func BenchmarkQuickSort(b *testing.B) {
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

func BenchmarkRadixSort(b *testing.B) {
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
