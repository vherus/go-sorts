package sort

func Merge(items []int) []int {
	length := len(items)

	if length == 1 {
		return items
	}

	halfWay := length / 2

	left, right := items[:halfWay], items[halfWay:]

	recursiveLeft, recursiveRight := Merge(left), Merge(right)

	return merge(recursiveLeft, recursiveRight)
}

func merge(left []int, right []int) []int {
	leftLength, rightLength := len(left), len(right)
	result := make([]int, leftLength+rightLength)

	lIndex, rIndex, resultIndex := 0, 0, 0

	for {
		if lIndex >= leftLength || rIndex >= rightLength {
			break
		}

		if left[lIndex] < right[rIndex] {
			result[resultIndex] = left[lIndex]
			resultIndex++
			lIndex++
			continue
		}

		result[resultIndex] = right[rIndex]
		resultIndex++
		rIndex++
		continue
	}

	// now either left or right will have been fully checked
	// so make sure to account for the remaining elements

	for {
		if lIndex < leftLength {
			result[resultIndex] = left[lIndex]
			resultIndex++
			lIndex++
			continue
		}

		break
	}

	for {
		if rIndex < rightLength {
			result[resultIndex] = right[rIndex]
			resultIndex++
			rIndex++
			continue
		}

		break
	}

	return result
}
