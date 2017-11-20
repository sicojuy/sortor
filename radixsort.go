package sortor

const (
	bits  = 8
	radix = 0xFF
)

var counter [radix + 1]int

func initCounter() {
	for i := 0; i <= radix; i++ {
		counter[i] = 0
	}
}

func maxBit(arrs []int) int {
	max := 0
	for _, x := range arrs {
		if max < x {
			max = x
		}
	}
	blen := 0
	for max > 0 {
		blen++
		max = max >> 1
	}
	return blen
}

func countSort(src []int, dst []int, offset uint) {
	initCounter()
	for _, x := range src {
		counter[x>>offset&radix]++
	}
	for i := 1; i <= radix; i++ {
		counter[i] += counter[i-1]
	}
	for i := len(src) - 1; i >= 0; i-- {
		j := src[i] >> offset & radix
		counter[j]--
		dst[counter[j]] = src[i]
	}
}

func RadixSort(arrs []int) {
	blen := maxBit(arrs)
	tmp := make([]int, len(arrs))
	src := &arrs
	dst := &tmp
	for i := 0; i*bits < blen; i++ {
		countSort(*src, *dst, uint(i*bits))
		src, dst = dst, src
	}
	if src != &arrs {
		for i := range arrs {
			arrs[i] = (*src)[i]
		}
	}
}
