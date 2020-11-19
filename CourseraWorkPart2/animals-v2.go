package main

import(

)

type Animals interface {
	Eat()
	Move()
	Speak()
}
type AnimalsType struct {
	Name string
	Food string
	Locomotion string
	Noise string
}
type AnimalsStorage struct {
	storage []AnimalsType
}

func main (){

}

func CreateAnimal(f string,l string,n string,nm string)AnimalsType{
	newAnimal := AnimalsType{
		Name: nm,
		Food:       f,
		Locomotion: l,
		Noise:      n,
	}
	return newAnimal
}

func StorageAnimals (animal []AnimalsType, newAnimal AnimalsType)[]AnimalsType{
	animal = append(animal,newAnimal)
	return animal
}

func Query (animal []AnimalsType, name string, command string){
	for _,key := range animal {
		if key.Name == name{
			if command == "speak"{

			}
			if command == "eat"{

			}
			if command == "move"{

			}
			return
		}
	}
	println("NOT Found: ",name)
}
func(a *AnimalsType) Eat(){

}