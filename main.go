package main

import (
	"fmt"
)

func main() {
	a := []int{5, 1, 2, 5}
	b := []int{4, 2, 5, 1, 1, 2}
	fmt.Print(backUniqueNumbers(a))
	fmt.Println(backUniqueNumbers(b))
	fmt.Println(intersection(a, b))
	fmt.Println(backUniqueNumbersTwoSlices(a, b))
}
func backUniqueNumbers(nums []int) []int {
	numbers := make(map[int]bool)
	var result []int
	for _, num := range nums {
		if !numbers[num] {
			numbers[num] = true
			result = append(result, num)

		}
	}
	return result
}
func intersection(nums1, nums2 []int) []int {
	numbers1 := make(map[int]bool)
	numbers2 := make(map[int]bool)
	var result []int
	for _, num := range nums1 {
		numbers1[num] = true
	}
	for _, num := range nums2 {
		numbers2[num] = true
	}
	for key := range numbers1 {
		if numbers2[key] {
			result = append(result, key)
		}
	}
	return result
}

func backUniqueNumbersTwoSlices(nums1, nums2 []int) []int {
	numbers := make(map[int]bool)
	var result []int
	for _, num := range nums1 {
		numbers[num] = true
	}
	for _, num := range nums2 {
		numbers[num] = true
	}
	for key := range numbers {

		result = append(result, key)

	}
	return result
}
