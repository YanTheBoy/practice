package week3_multithreaing

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

func CmdLineReader() {
	f, _ := os.Create("file.txt")
	defer f.Close()

	readChan := make(chan string)
	sigchan := make(chan os.Signal)
	mu := new(sync.Mutex)

	go func() {
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Считывание остановлено.")

		err := f.Close()
		if err != nil {
			log.Printf("Ошибка при чтении строки %s\n", err)
		}
		close(readChan)
		os.Exit(0)
	}()

	var a string
	go FileWriter(readChan, f, mu)
	for {
		_, err := fmt.Scan(&a)
		if err != nil {
			log.Printf("Ошибка при чтении строки %s\n", err)
		}
		readChan <- a
	}

}

func FileWriter(ch chan string, f *os.File, mu *sync.Mutex) {
	for word := range ch {
		mu.Lock()
		_, err := f.WriteString(word + "\n")
		if err != nil {
			log.Printf("Ошибка при записи строки %s\n", err)
		}
		mu.Unlock()
	}

}


/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

