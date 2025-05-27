package main

import (
	presenter "Spam-Masker/basic/presenters"
	producer "Spam-Masker/basic/producers"
	service "Spam-Masker/basic/services"
	"fmt"
)

func main() {
	produser := &producer.FileProducer{FilePath: "text.txt"}
	presenter := &presenter.FilePresenter{FilePath: "mask_text.txt"}
	serv := service.NewService(produser, presenter)
	if err := serv.Run(); err != nil {
		fmt.Println("Ошибка Run()", err)
	}
}
