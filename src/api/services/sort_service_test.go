package services

import (
	"fmt"
	"testing"
)

func print(s ...interface{}) {
	fmt.Println("          ->", s)
}

func TestPrivateConst(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}

func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}

	print("before:", elements)
	Sort(elements)
	print("after:", elements)

	if elements[0] != 0 {
		t.Error("First element should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("Last element should be 9")
	}

}
