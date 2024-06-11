package week2_maps

func IsEqualArrays[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int)
	for _, item := range a {
		counts[item]++
	}

	for _, item := range b {
		if counts[item] == 0 {
			return false
		}
		counts[item]--
	}

	return true
}

/* 8. Напишите функцию-дженерик IsEqualArrays для comparable типов, которая сравнивает два неотсортированных массива.
Функция выдает булевое значение как результат. true - если массивы равны, false - если нет.
Массивы считаются равными, если в элемент из первого массива существует в другом, и наоборот.
Вне зависимости от расположения.
*/
