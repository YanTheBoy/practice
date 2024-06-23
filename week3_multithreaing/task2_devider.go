package week3_multithreaing

import (
	"math"
)

func SieveOfEratosthenes(m []int) (primes, composites []int) {
	var p, c []int

	for x := range FindPrimes(m) {
		p = append(p, x)
	}
	for x := range FindComposites(m) {
		c = append(c, x)
	}
	return p, c
}

func FindPrimes(m []int) chan int {
	var ch = make(chan int)
	go func() {
		defer close(ch)
		for _, val := range m {
			if isPrimeNumber(val) {
				ch <- val
			}
		}
	}()
	return ch
}

func FindComposites(m []int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, val := range m {
			if !isPrimeNumber(val) {
				ch <- val
			}
		}
	}()
	return ch
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

