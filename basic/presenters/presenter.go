package presenter

import (
	"fmt"
	"os"
	"strings"
)

type FilePresenter struct {
	FilePath string
}

func (f *FilePresenter) Present(p []string) error {
	file, err := os.Create(f.FilePath)
	if err != nil {
		fmt.Println("Ошибка создания файла")
		return err
	}
	defer file.Close()
	_, err = file.WriteString(strings.Join(p, " "))
	if err != nil {
		fmt.Println("Ошибка записи в файл")
		return err
	}
	return nil
}
