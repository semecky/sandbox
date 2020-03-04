package ducktyping

import "fmt"

type MeatEater struct {
}
type Veggie struct {
}

func (v Veggie) EatVegetable() {
	fmt.Printf("i am %T and i am eating vegetables\n", v)
}
func (m MeatEater) EatMeat() {
	fmt.Printf("i am %T and i am eating meat\n", m)
}

//Человек, ест мясо и овощи
type Human struct {
	MeatEater
	Veggie
}

//Овца, есть овощи
type Sheep struct {
	Veggie
}

//Волк, ест мясо
type Wolf struct {
	MeatEater
}

//Травоядное, ест только овощи
type Herbivore interface {
	EatVegetable()
}

//Плотоядное, есть только мясо
type Carnivore interface {
	EatMeat()
}

//Всеядное, может есть мясо и овощи
type Omnivore interface {
	EatVegetable()
	EatMeat()
}
