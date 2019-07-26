package sortor

func HeapSort(arrs []int) {
	size := len(arrs)
	for i := 0; i < size; i++ {
		j := i
		for j > 0 {
			k := (j - 1) >> 1
			if arrs[j] < arrs[k] {
				break
			}
			arrs[j], arrs[k] = arrs[k], arrs[j]
			j = k
		}
	}
	for i := size - 1; i > 0; i-- {
		arrs[i], arrs[0] = arrs[0], arrs[i]
		j := 0
		for (j<<1)+1 < i {
			k := (j << 1) + 1
			if (j+1)<<1 < i && arrs[k] < arrs[(j+1)<<1] {
				k = (j + 1) << 1
			}
			if arrs[j] < arrs[k] {
				arrs[j], arrs[k] = arrs[k], arrs[j]
				j = k
			} else {
				break
			}
		}
	}
}
