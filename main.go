package main

import (
	"fmt"

	"github.com/vherus/go-sorts/sort"
)

func main() {
	unsorted := []int{12, 7, 923, 64, 198, 71, 28, 88}
	unsortedMsg := fmt.Sprintf("Unsorted: %v", unsorted)
	fmt.Println(unsortedMsg)

	selective := sort.Selective(unsorted)
	selectiveMsg := fmt.Sprintf("Selective: %v", selective)
	fmt.Println(selectiveMsg)

	merge := sort.Merge(unsorted)
	mergeMsg := fmt.Sprintf("Merge: %v", merge)
	fmt.Println(mergeMsg)
}
