package main

import (
	"errors"
	"fmt"
)

// ================ ENUMS =================

type Category int

const (
	Mammal Category = iota
	Reptile
	Avian
)

// ============================================

// =============== INTERFACES ==============

type Dangerous interface {
	DangerLevel() string
}

type ZooAnimal interface {
	Sound() string
	Move() string
	Info() string
	Name() string
}

// ============================================

// =============== STRUCTS ==============

type Animal struct {
	name     string
	age      uint16
	weight   float32
	category Category
}

func (a Animal) Name() string {
	return a.name;
}

type Dog struct {
	Animal
	breed string
}

type Bird struct {
	Animal
	canFly bool
}

type Snake struct {
	Animal
	venomous bool
}

// ============================================

// =============== METHODS ==============

func (dog Dog) Sound() string {
	return "Woof woof!"
}

func (dog Dog) Move() string {
	return "Running on four legs..."
}

func (dog Dog) Info() string {
	baseInfo := baseAnimalInfo(dog.Animal)
	return fmt.Sprintf(`%v
breed-> %v`, baseInfo, dog.breed)
}

func (bird Bird) Sound() string {
	return "Chirp chirp!"
}

func (bird Bird) Move() string {
	return "Flying through the sky..."
}

func (bird Bird) Info() string {
	baseInfo := baseAnimalInfo(bird.Animal)
	return fmt.Sprintf(`%v
Can Fly-> %v`, baseInfo, bird.canFly)
}

func (snake Snake) Sound() string {
	return "Ssssss..."
}

func (snake Snake) Move() string {
	return "Slithering on the ground..."
}

func (snake Snake) Info() string {
	baseInfo := baseAnimalInfo(snake.Animal)
	return fmt.Sprintf(`%v
Venomous -> %v`, baseInfo, snake.venomous)
}

func (snake Snake) DangerLevel() string {
	return "Moderate"
}

// =============== UTILS ==============

var speciesNames = map[Category]string{
	Mammal:  "mammal",
	Reptile: "reptile",
	Avian:   "avian",
}

func (cat Category) String() string {
	return speciesNames[cat]
}

func baseAnimalInfo(a Animal) string {
	return fmt.Sprintf(`name -> %v
Age -> %v
Category -> %v
weight -> %vkg`, a.name, a.age, a.category, a.weight)
}

func NewAnimal(name string, age uint16, weight float32, species Category) Animal {
	return Animal{
		name:     name,
		age:      age,
		weight:   weight,
		category: species,
	}
}

func NewDog(name string, age uint16, weight float32, breed string) Dog {
	return Dog{
		Animal: NewAnimal(name, age, weight, Mammal),
		breed:  breed,
	}
}

func NewBird(name string, age uint16, weight float32, canFly bool) Bird {
	return Bird{
		Animal: NewAnimal(name, age, weight, Avian),
		canFly: canFly,
	}
}

func NewSnake(name string, age uint16, weight float32, venemous bool) Snake {
	return Snake{
		Animal:   NewAnimal(name, age, weight, Reptile),
		venomous: venemous,
	}
}

var AnimalNotFoundError = errors.New("Animal was not found with the given name")

func FindByName(animals []ZooAnimal, name string) (ZooAnimal, error) {
	for _, a := range animals {
		if(a.Name() == name) {
			return a, nil;
		}
	}

	return nil, AnimalNotFoundError;
}

func reportAnimals(animals []ZooAnimal) {
	for i, animal := range animals {
		fmt.Printf("====== Animal %v ======\n", i+1)
		fmt.Println(animal.Info())
		if dangerous, ok := animal.(Dangerous); ok {
			fmt.Println("DANGEROUS ANIMAL. Dangerous Level -> ", dangerous.DangerLevel())
		}
	}
}

func FilterAnimal[T any](animals []T, predicate func(T) bool) []T {
	result := []T{}

	for _, a := range animals {
		if predicate(a) {
			result = append(result, a)
		}
	}

	return result
}

func main() {
	animals := []ZooAnimal{
		NewDog("Churros", 11, 8, "breed1"),
		NewDog("Shoyou", 11, 8, "breed2"),
		NewBird("zeca", 15, 3, true),
		NewBird("loro", 10, 2, true),
		NewSnake("cobra", 8, 2, false),
		NewSnake("king", 10, 2, true),
	}
	reportAnimals(animals)

	snakes := FilterAnimal(animals, func(animal ZooAnimal) bool {
		_, ok := animal.(Snake)
		return ok
	})

	fmt.Println()
	fmt.Println("Reporting only snakes")
	reportAnimals(snakes)

	animal, _:= FindByName(animals, "Churros");
	fmt.Println("Animal:", animal.Info());
}
