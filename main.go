package main

import (
	"fmt"
	"task8.8.1/electronic"
)

func main() {
	var app = electronic.NewApplePhone("IPHONE 6", "IOS6")
	var andr = electronic.NewandroidPhone("Xiaomi", "Redmi Note 8", "Android 11")
	var station = electronic.NewradioPhone("Motorola", "C1001CB", "20")
	printCharacteristics(app)
	printCharacteristics(andr)
	printCharacteristics(station)
}

func printCharacteristics(e electronic.Phone) {
	fmt.Printf("Брэнд: %s, Модель: %s, Тип: %s, ", e.Brand(), e.Model(), e.Type())
	if s, ok := e.(electronic.Smartphone); ok {
		fmt.Println("ОС:", s.OS())
	}
	if r, ok := e.(electronic.StationPhone); ok {
		fmt.Println("Количество кнопок:", r.ButtonsCount())
	}
}

//	Newapplephone := electronic.NewApplePhone("fdgfd", "")
//	fmt.Println(Newapplephone)
//
//}
