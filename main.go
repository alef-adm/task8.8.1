package main

import (
	"fmt"
	"task8.8.1/electronic"
)

func main() {
	var app = electronic.NewApplePhone("IPHONE 6", "IOS6")
	var andr = electronic.NewandroidPhone("Xiaomi", "Redmi Note 8", "Android 11")
	printCharacteristics(app)
	printCharacteristics(andr)
}

func printCharacteristics(e electronic.Phone) {
	fmt.Printf("Брэнд: %s, Модель: %s, Тип: %s, ", e.Brand(), e.Model(), e.Type())
	if s, ok := e.(electronic.Smartphone); ok {
		fmt.Println("ОС:", s.OS())
	}
}

//	Newapplephone := electronic.NewApplePhone("fdgfd", "")
//	fmt.Println(Newapplephone)
//
//}
