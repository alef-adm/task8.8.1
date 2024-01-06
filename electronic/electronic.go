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
	ButtonsCount() string
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

type radioPhone struct {
	brand string
	model string
	cob   string
}

func (c *radioPhone) Brand() string {
	return c.brand
}
func (c *radioPhone) Model() string {
	return c.model
}
func (c *radioPhone) Type() string {
	return "station"
}
func (c *radioPhone) ButtonsCount() string {
	return c.cob
}
func NewradioPhone(brand, model, cob string) *radioPhone {
	return &radioPhone{brand: brand, model: model, cob: cob}
}

//func NewapplePhone(model, typephone string, weight int, diagonal float64) applePhone {
//	return applePhone{model: Phone.Model(), typephone: Phone.Type(), os: Smartphone.OS(), weight: weight, diagonal: diagonal}
