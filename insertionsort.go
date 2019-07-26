package sortor

func InsertionSort(arrs []int) {
	size := len(arrs)
	for i := 0; i < size; i++ {
		for j := i; j > 0; j-- {
			if arrs[j] < arrs[j-1] {
				arrs[j], arrs[j-1] = arrs[j-1], arrs[j]
			}
		}
	}
}
