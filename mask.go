package main

import (
	"bufio"
	"fmt"
	"os"
)

func ScanMask() string { // Функция сканирует текст из буфера и возвращает введеный текст с замаскированным URL
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите текст: ")
	input, _ := reader.ReadString('\n')
	var (
		sTOb     []byte = []byte(input)     // Переводит вводимую строку в срез байтов
		bHTTP    []byte = []byte("http://") // Задаем (http://) виде байтового среза
		sizeHTTP int    = len(bHTTP)        // Размер bHTTP
		size     int    = len(sTOb)
		output   []byte = make([]byte, 0, size)
	)

	// Итерируемся по массиву

	for i := 0; i < size; i++ {

		// Находит положение http://

		if i+sizeHTTP <= size && string(sTOb[i:(i+sizeHTTP)]) == "http://" {
			output = append(output, sTOb[i:i+sizeHTTP]...) // Начало + http://
			i += sizeHTTP

			// Заменяет все после http:// на * до пробела

			for i < size && sTOb[i] != ' ' {
				output = append(output, '*')
				i++
			}

			if i < size {
				output = append(output, sTOb[i])
			}
			continue
		}
		output = append(output, sTOb[i])

	}
	return string(output)
}

func main() {
	var output string = ScanMask()
	var outPut string = output
	fmt.Println(outPut)
}
