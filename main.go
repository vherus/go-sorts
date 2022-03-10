package main

import (
	"fmt"

	"github.com/vherus/go-sorts/sort"
)

func main() {
	unsorted := []int{12, 7, 923, 64, 198, 71, 28}

	selective := sort.Selective(unsorted)
	selectiveMsg := fmt.Sprintf("Selective: %v", selective)
	fmt.Println(selectiveMsg)
}
