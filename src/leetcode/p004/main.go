// go:build 004
package main

import (
	"fmt"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0
	}

	aim := append(nums1, nums2...)
	sort.Ints(aim)

	if len(aim)%2 == 0 {
		return float64(aim[len(aim)/2-1]+aim[len(aim)/2]) / 2
	}

	return float64(aim[len(aim)/2])
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2, 4}))
}
