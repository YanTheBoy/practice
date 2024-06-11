package week2_maps

import (
	s "slices"
)

/* 1. Напишите функцию, которая находит пересечение неопределенного количества слайсов типа int.
Каждый элемент в пересечении должен быть уникальным. Слайс-результат должен быть отсортирован в восходящем порядке.
Примеры:
1. Если на вход подается только 1 слайс [1, 2, 3, 2], результатом должен быть слайс [1, 2, 3].
2. Вход: 2 слайса [1, 2, 3, 2] и [3, 2], результат - [2, 3].
3. Вход: 3 слайса [1, 2, 3, 2], [3, 2] и [], результат - [].
*/

func FindIntersections(slices ...[]int) []int {

	result := make(map[int]int)
	for _, slice := range slices {
		if len(slice) == 0 {
			return []int{}
		}

		tempMap := make(map[int]bool)
		for _, value := range slice {
			if !tempMap[value] {
				tempMap[value] = true
				result[value]++
			}
		}
	}

	var crossed []int
	for value, count := range result {
		if count == len(slices) {
			crossed = append(crossed, value)
		}
	}

	s.Sort(crossed)
	return crossed
}
