package pattern

import (
	"fmt"
)

/*
	Реализовать паттерн «фабричный метод».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
type Product interface {
	someMethod()
}

type ConcreteProduct struct {
	name string
}

func (c *ConcreteProduct) someMethod() {
	fmt.Println(c.name)
}

type ConcreteProductA struct {
	ConcreteProduct
}

func NewConcreteProductA(name string) *ConcreteProductA {
	concreteProductA := new(ConcreteProductA)
	concreteProductA.name = name
	return concreteProductA
}

type ConcreteProductB struct {
	ConcreteProduct
}

func NewConcreteProductB(name string) *ConcreteProductB {
	concreteProductB := new(ConcreteProductB)
	concreteProductB.name = name
	return concreteProductB
}

func GetProduct(productType string, name string) (Product, error) {
	switch productType {
	case "ConcreteProductA":
		return NewConcreteProductA(name), nil
	case "ConcreteProductB":
		return NewConcreteProductB(name), nil
	default:
		return nil, fmt.Errorf("не правильный тип продукта")
	}
}
