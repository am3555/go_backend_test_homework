package main

import (
	"fmt"
	"sort" // или "slices"
)

// Median возвращает медиану числовой последовательности.
// Напишите код функции
// ...

func Median(list []int) int {
	sort.Ints(list)
	if len(list) == 0 {
		return 0
	}
	var med int
	if len(list)%2 == 0 {
		index := len(list) / 2
		med = (list[index-1] + list[index]) / 2
	} else {
		index2 := len(list) / 2
		med = (list[index2])
	}
	return med
}

func main() {
	lists := [][]int{
		{},
		{57},
		{78, -7},
		{99, 200, 0},
		{4, 4, 4, 3},
		{102, -7, 44, -7, 102},
		{82, -23, 1, 5, 98, 100},
		{100000, 90000, 20000, 20000, 20000, 22000, 25500, 22000},
	}
	medians := []int{
		0, 57, 35, 99, 4, 44, 43, 22000,
	}

	for i, list := range lists {
		if median := Median(list); median != medians[i] {
			fmt.Printf("median %d: %d != %d\n", i, medians[i], median)
		}
	}
	fmt.Println("Тестирование завершено")
}
