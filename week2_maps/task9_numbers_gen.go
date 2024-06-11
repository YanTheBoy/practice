package week2_maps

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type Numbers[T constraints.Integer] []T

func (n *Numbers[T]) Sum() {
	var calcSum T
	for _, num := range *n {
		calcSum += num
	}
	fmt.Println(calcSum)
}

func (n *Numbers[T]) Mult() {
	var calcMult T
	for _, num := range *n {
		calcMult += num
	}
	fmt.Println(calcMult)
}

func (n *Numbers[T]) IsEqual(s []T) {
	fmt.Println(len(*n) == len(s))
}

func (n *Numbers[T]) IsIncluded(num T) {
	for pos, val := range *n {
		if num == val {
			fmt.Println(pos)
			break
		}
	}
}

func (n Numbers[T]) PopByValue(num T) {
	for pos, val := range n {
		if val == num {
			n = append(n[:pos], n[pos+1:]...)
			break
		}
	}
	fmt.Println(n)
}

func (n Numbers[T]) PopByIndex(ind T) {
	n = append(n[:ind], n[ind+1:]...)
	fmt.Println(n)
}

/* 9. Реализуйте тип-дженерик Numbers, который является слайсом численных типов.
Реализуйте следующие методы для этого типа:
* суммирование всех элементов,
* произведение всех элементов,
* сравнение с другим слайсом на равность,
* проверка аргумента, является ли он элементом массива, если да - вывести индекс первого найденного элемента,
* удаление элемента массива по значению,
* удаление элемента массива по индексу.
*/
