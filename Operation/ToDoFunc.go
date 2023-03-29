package Operation

import (
	"accessProject/InputOutput"
	"fmt"
)

func ToDoFunc() {
	InputOutput.IntOutData()

	var firstDigit = InputOutput.Num / 10000
	var NumFloor = InputOutput.FloorNum

	switch firstDigit {
	case 1:
		if firstDigit == NumFloor {
			fmt.Println("У вас есть доступ в первый этаж")

		} else {
			fmt.Println("У вас нет доступа на данный этаж")
		}

	case 2:
		if firstDigit == NumFloor || firstDigit > NumFloor {
			fmt.Println("У вас есть доступ на данный этаж")

		} else {
			fmt.Println("У вас нет доступа на данный этаж")
		}

	case 3:
		if firstDigit == NumFloor || firstDigit > NumFloor {
			fmt.Printf("У вас есть доступ во все этажи здания, включая %v", NumFloor)

		} else {
			fmt.Println("У вас нет доступа на данный этаж")
		}

	default:
		fmt.Println("В доступе отказано")
		return
	}
}
