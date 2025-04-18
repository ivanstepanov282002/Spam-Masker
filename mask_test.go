package main

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanMask(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Текст без URL",
			input:    "Обычный текст\n",
			expected: "Обычный текст\n",
		},
		{
			name:     "Текст с URL в середине",
			input:    "Сайт http://example.com текст\n",
			expected: "Сайт http://*********** текст\n", // 11 звёздочек для "example.com" (11 символов)
		},
		{
			name:     "Только URL",
			input:    "http://site.ru\n",
			expected: "http://*******\n", // 7 звёздочек для "site.ru" (7 символов)
		},
		{
			name:     "URL с поддоменами",
			input:    "Адрес http://sub.site.ru\n",
			expected: "Адрес http://***********\n", // 11 звёздочек для "sub.site.ru" (11 символов)
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reader := bufio.NewReader(strings.NewReader(tt.input))
			result := mockScanMask(reader)

			if result != tt.expected {
				t.Errorf("Тест '%s' не пройден:\nожидалось: % q\nполучили:  % q",
					tt.name, tt.expected, result)
			}
		})
	}
}

func mockScanMask(reader *bufio.Reader) string {
	// Полная копия вашей исходной функции ScanMask
	input, _ := reader.ReadString('\n')
	var (
		sTOb     = []byte(input)
		bHTTP    = []byte("http://")
		sizeHTTP = len(bHTTP)
		size     = len(sTOb)
		output   = make([]byte, 0, size)
	)

	for i := 0; i < size; i++ {
		if i+sizeHTTP <= size && string(sTOb[i:(i+sizeHTTP)]) == "http://" {
			output = append(output, sTOb[i:i+sizeHTTP]...)
			i += sizeHTTP

			for i < size && sTOb[i] != ' ' && sTOb[i] != '\n' {
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
