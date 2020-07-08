package sort

func BubbleSort(a []int) {
	l := len(a)
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}
