package sortor

func SelectionSort(arrs []int) {
	size := len(arrs)
	for i := 0; i < size; i++ {
		k := i
		for j := i + 1; j < size; j++ {
			if arrs[j] < arrs[k] {
				k = j
			}
		}
		if i != k {
			arrs[i], arrs[k] = arrs[k], arrs[i]
		}
	}
}
