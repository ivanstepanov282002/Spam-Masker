package producer

import (
	"fmt"
	"os"
	"strings"
)

type FileProducer struct {
	FilePath string
}

func (p *FileProducer) Produce() ([]string, error) {
	slicesByteText, err := os.ReadFile(p.FilePath)
	if err != nil {
		fmt.Println("Ошибка чтения файла")
		return nil, err
	}
	slicesText := strings.Split(string(slicesByteText), "/n")
	return slicesText, nil
}
