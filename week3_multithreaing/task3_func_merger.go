package week3_multithreaing

import (
	"sync"
)

func Merger(channels... chan int) <-chan int {
	outCh := make(chan int)
	var wg sync.WaitGroup
	//wg.Add(len(channels))
	for _, ch := range channels {
		wg.Add(1)
		go func(ch chan int) {
			defer wg.Done()
			for v := range ch {
				outCh <- v
			}
		}(ch)

	}
	go func() {
		wg.Wait()
		close(outCh)
	}()
	return outCh
}


/*3. Реализуйте функцию слияния двух каналов в один.*/
