package main

import (
	"strings"
)

func AddFromArray(array []int) ([]int, error) {

	sum := 0
	for _, v := range array {
		sum += v
	}

	return []int{sum}, nil
}

func Max(array []int) (int, error) {

	if len(array) == 0 {
		panic("Array darf nicht leer sein")
	}

	maxNum := array[0]

	for _, i := range array {
		if i > maxNum {
			maxNum = i
		}
	}
	return maxNum, nil
}

func Min(array []int) (int, error) {
	if len(array) == 0 {
		panic("Array darf nicht leer sein")
	}

	min := array[0]

	for _, i := range array {
		if i < min {
			min = i
		}
	}

	return min, nil
}

func WordCounter(text string) (map[string]int, error) {

	words := strings.Fields(text)

	counts := make(map[string]int)

	for _, w := range words {
		counts[w]++
	}

	return counts, nil
}
