package actioninfo

import "log"

type DataParser interface {
	// TODO: добавить методы
	Parse(string) error
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	// TODO: реализовать функцию
	for _, data := range dataset {
		err := dp.Parse(data)
		if err != nil {
			log.Fatalf("ошибка парсинга данных: %v", err)
			continue
		}
		info, err := dp.ActionInfo()
		if err != nil {
			log.Fatalf("ошибка получения информации о действии: %v", err)
			continue
		}
		log.Printf("Информация о действии: %s", info)
	}
}
