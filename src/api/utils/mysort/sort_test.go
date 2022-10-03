package mysort

import (
	"fmt"
	"testing"
)

func TestBubbleSortOrderDESC(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}

	fmt.Println(elements)
	BubbleSort(elements)
	fmt.Println(elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}

}
