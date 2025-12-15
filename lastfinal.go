package main

import (
    "fmt"
    "sort" // или "slices"
)
// map[57:1]
// map[-7:1 78:1]
// map[0:1 99:1 200:1]
// map[3:1 4:3]
// map[-7:2 44:1 102:2]
// map[-23:1 1:1 5:1 82:1 98:1 100:1]
// map[20000:3 22000:2 25500:1 90000:1 100000:1]
// Mode возвращает моды числовой последовательности.
// Напишите код функции
// ...
func Mode(numbers []int) ([]int, int) {
	if len(numbers) == 0 {
		return []int {}, 1
	}

	mapa := make(map[int]int)
	for _, value := range numbers {
		mapa[value]++
		fmt.Println(len(numbers), len(mapa))
		if len(numbers) == len(mapa) {
			return []int {}, 1
		}
	}

	var maxval int
	for _, value := range mapa {
		if value > maxval {
			maxval = value
		}
	}

	resultSlice := make([]int, 0)
	for k, v := range mapa {
		if v == maxval {
			resultSlice = append(resultSlice, k)
		}
	}

	sort.Ints(resultSlice)
	return resultSlice, maxval

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
    modes := [][]int{
        {}, {}, {}, {},
        {4},
        {-7, 102}, {},
        {20000},
    }
    mcount := []int{
        1, 1, 1, 1, 3, 2, 1, 3,
    }
    for i, list := range lists {
        mode, count := Mode(list)
        if len(mode) != len(modes[i]) {
            fmt.Printf("len mode %d: %v != %v'\n", i, modes[i], mode)
        } else {
            for j, v := range mode {
                if v != modes[i][j] {
                   fmt.Printf("mode %d: %v != %v\n", i, modes[i], mode)
                }
            }
        }
        if count != mcount[i] {
            fmt.Printf("mcount %d: %d != %d\n", i, mcount[i], count)
        }
    }
    fmt.Println("Тестирование завершено")
}

// Для сортировки слайсов в функциях Mode() используйте функцию sort.Ints() или slices.Sort().
// Вам нужно будет создать map[int]int и за один проход подсчитать, сколько раз встречаются числа в последовательности. 
// Затем переберите мапу и определите максимальное количество вхождений. 
// Если оно больше 1, значит, как минимум одна мода есть. 
// Остаётся ещё раз пройтись по мапе и добавить в результирующий слайс все числа с максимальным количеством вхождений.