package services

import (
	"testing"
)

func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}

	Sort(elements)

	if elements[0] != 9 {
		t.Error("First element should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("Last element should be 0")
	}

}
