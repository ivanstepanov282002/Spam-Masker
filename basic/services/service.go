package service

type Producer interface {
	Produce() ([]string, error)
}

type Presenter interface {
	Present([]string) error
}

type Service struct {
	prod Producer
	pres Presenter
}

func NewService(p Producer, pr Presenter) *Service {
	return &Service{prod: p, pres: pr}
}

func (s *Service) Mask(pr string) string {
	var (
		sTOb     []byte = []byte(pr)        // Переводит вводимую строку в срез байтов
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

func (s *Service) Run() error {
	slText, err := s.prod.Produce()
	if err != nil {
		return err
	}

	sL := make([]string, 0, len(slText))

	for _, v := range slText {
		sL = append(sL, s.Mask(v))
	}

	if err := s.pres.Present(sL); err != nil {
		return err
	}
	return nil
}
