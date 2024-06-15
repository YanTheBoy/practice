package week3_multithreaing

import (
	"fmt"
	"log"
	"os"
	"os/signal"
)

func CmdLineReader() []string {
	var list []string
	var a string

	go func() {
		sigchan := make(chan os.Signal)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		log.Println("Считывание остановлено.")
		FileWriter(list)
		os.Exit(0)
	}()
	for {
		fmt.Scan(&a)
		list = append(list, a)
	}

	return list
}

func FileWriter(data []string) {

	f, _ := os.Create("file.txt")
	defer f.Close()

	for _, i := range data {
		f.WriteString(i + "\n")
	}

}

/*1. Напишите 2 функции:
	Первая функция читает ввод с консоли. Ввод одного значения заканчивается нажатием клавиши enter.
	Вторая функция пишет эти данные в файл. Свяжите эти функции каналом.
Работа приложения должна завершится при нажатии клавиш ctrl+c с кодом 0. */

/*2. Напишите функцию разделения массива чисел на массивы простых и составных чисел.
Для записи в массивы используйте два разных канала и горутины.
Важно, чтобы были использованы владельцы каналов.*/

/*3. Реализуйте функцию слияния двух каналов в один.*/
