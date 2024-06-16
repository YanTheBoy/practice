package week3_multithreaing

import (
	"fmt"
	"math"
	"slices"
)

func Devider(m []int) {

}

func SieveOfEratosthenes(m []int) {
	slices.Sort(m)
	var primes, composites []int
	for _, val := range m {
		if isPrimeNumber(val) {
			primes = append(primes, val)
		} else {
			composites = append(composites, val)
		}
	}
	fmt.Println(primes)
	fmt.Println(composites)

}

func isPrimeNumber(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

/*2. Напишите функцию разделения массива чисел на массивы простых и составных чисел.
Для записи в массивы используйте два разных канала и горутины.
Важно, чтобы были использованы владельцы каналов.*/

/*3. Реализуйте функцию слияния двух каналов в один.*/
