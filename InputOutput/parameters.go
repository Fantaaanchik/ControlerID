package InputOutput

import (
	"fmt"
	"log"
)

var (
	Num      int32
	id       int32
	floor    int32
	FloorNum int32
)

func IntOutData() {

	fmt.Printf("Вход строго по пропускам.\nВведите ваш ID:\n")
	_, err := fmt.Scan(&id)
	if err != nil {
		log.Println("Вы ввели не правильный формат данных, нужно ввести число\n", err.Error())
		return
	}

	if id > 39999 || id < 10000 {
		fmt.Printf("Вы ввели не правильный ID, Нужно ввести число в диапазоне 10000..39999\n")
		return
	}

	fmt.Printf("Введите этаж на который хотите попасть:\n")
	_, er := fmt.Scan(&floor)
	if er != nil {
		log.Println("Вы ввели не правильный формат данных, нужно ввести число\n", er.Error())
		return
	}

	Num = id
	FloorNum = floor
	return
}
