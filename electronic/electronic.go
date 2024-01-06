package electronic

import (
	"fmt"
)

type Phone interface {
	Brand() string
	Model() string
	Type() string
}

type StationPhone interface {
	ButtonsCount(int) int
}

type Smartphone interface {
	OS() string //название операционной системы
}

// applephone
type applePhone struct {
	m string
	o string
}

func (a *applePhone) Brand() string {
	return "Apple"
}
func (a *applePhone) Model() string {
	return a.m
}
func (a *applePhone) Type() string {
	return fmt.Sprintln("SmartPhone")
}
func (a *applePhone) OS() string {
	return a.o
}

func NewApplePhone(model, os string) *applePhone {
	return &applePhone{m: model, o: os}
}

type androidPhone struct {
	brand string
	model string
	os    string
}

func (b *androidPhone) Brand() string {
	return b.brand
}
func (b *androidPhone) Model() string {
	return b.model
}
func (b *androidPhone) Type() string {
	return "SmartPhone"
}
func (b *androidPhone) OS() string {
	return b.os
}
func NewandroidPhone(brand, model, os string) *androidPhone {
	return &androidPhone{brand: brand, model: model, os: os}
}

//
//type radioPhone struct {
//}
//
//func (c *radioPhone) Brand(brand string) string {
//	return brand
//}
//func (c *radioPhone) Model(model string) string {
//	return model
//}
//func (c *radioPhone) Type() string {
//	return "station"
//}
//func (c *radioPhone) ButtonsCount(countOfButton int) int {
//	return countOfButton
//
//}

//func NewapplePhone(model, typephone string, weight int, diagonal float64) applePhone {
//	return applePhone{model: Phone.Model(), typephone: Phone.Type(), os: Smartphone.OS(), weight: weight, diagonal: diagonal}
