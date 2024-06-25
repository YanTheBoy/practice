// FIX me pls

package main

import (
	"fmt"
	"maps"
	"sync"
	"time"
)


func RunProcessor(wg *sync.WaitGroup, prices []map[string]float64, mu *sync.Mutex) {
	defer wg.Done()
	func() {
     for _, price := range prices {
		 mu.Lock()
         for key, value := range price {
             price[key] = value + 1
         }
         fmt.Println(price)
		 mu.Unlock()
     }
 }()
}

func RunWriter(mu *sync.Mutex) <-chan map[string]float64 {
	var prices = make(chan map[string]float64)
	go func() {
		var currentPrice = map[string]float64{
			"inst1": 1.1,
			"inst2": 2.1,
			"inst3": 3.1,
			"inst4": 4.1,
		}


		for i := 1; i < 5; i++ {
			mu.Lock()
			currentPrice = maps.Clone(currentPrice)

			for key, value := range currentPrice {
				currentPrice[key] = value + 1
			}
			prices <- currentPrice
			mu.Unlock()
			time.Sleep(time.Second)
		}

		close(prices)
	}()
	return prices
}
func main() {
	var mu sync.Mutex

	p := RunWriter(&mu)
	var prices []map[string]float64

	for price := range p {
		prices = append(prices, price)
	}

	for _, price := range prices {
		fmt.Println(price)
	}

	wg := sync.WaitGroup{}
	wg.Add(3)
	go RunProcessor(&wg, prices, &mu)
	go RunProcessor(&wg, prices, &mu)
	go RunProcessor(&wg, prices, &mu)
	wg.Wait()
}
