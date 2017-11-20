package sortor

func insertSort(arrs []int) {
	size := len(arrs)
	if size <= 1 {
		return
	} else if size == 2 {
		if arrs[0] > arrs[1] {
			arrs[0], arrs[1] = arrs[1], arrs[0]
		}
		return
	} else if size == 3 {
		if arrs[0] > arrs[1] {
			arrs[0], arrs[1] = arrs[1], arrs[0]
		}
		if arrs[0] > arrs[2] {
			arrs[0], arrs[2] = arrs[2], arrs[0]
		}
		if arrs[1] > arrs[2] {
			arrs[1], arrs[2] = arrs[2], arrs[1]
		}
		return
	}
}

func pardition(arrs []int) int {
	size := len(arrs)
	i := 0
	for j := 1; j < size; j++ {
		if arrs[j] < arrs[i] {
			arrs[i+1], arrs[j] = arrs[j], arrs[i+1]
			arrs[i], arrs[i+1] = arrs[i+1], arrs[i]
			i++
		}
	}
	return i
}

func QuickSort(arrs []int) {
	if len(arrs) <= 3 {
		insertSort(arrs)
		return
	}
	i := pardition(arrs)
	if i > 0 {
		QuickSort(arrs[0:i])
	}
	if i < len(arrs) {
		QuickSort(arrs[i+1 : len(arrs)])
	}
}
