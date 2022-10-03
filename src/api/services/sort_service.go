package services

import (
	"github.com/zxhoper/go-testing-sort-integration/src/api/utils/mysort"
)

const (
	privateConst = "private"
)

func Sort(elements []int) {
	mysort.BubbleSort(elements)
}
