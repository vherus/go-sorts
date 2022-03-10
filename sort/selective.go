package sort

func Selective(items []int) []int {
	length := len(items)

	for i := 0; i < length; i++ {
		smallestIndex := i
		largestIndex := i

		for j := i + 1; j < length; j++ {
			if items[j] < items[smallestIndex] {
				smallestIndex = j
			}

			if items[j] > items[largestIndex] {
				largestIndex = j
			}
		}

		original := items[i]
		replacement := items[smallestIndex]

		items[i] = replacement
		items[smallestIndex] = original
	}

	return items
}
