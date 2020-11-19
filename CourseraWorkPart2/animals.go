package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	Food string
	Locomotion string
	Noise string
}

func main (){
	fmt.Println("Input The Type of Animal (cow/bird/snake) and Request (eat, move,speak)")
	reader :=bufio.NewReader(os.Stdin)
	fmt.Print("> ")
	readInput,err := reader.ReadString('\n')
	if err !=nil{
		fmt.Println("Read In Error")
		os.Exit(1)
	}
	animalInfo := strings.Split(strings.TrimSpace(readInput)," ")

	mal := new(Animal)

	if mal == nil {
		fmt.Println("Animal Not Found")
		os.Exit(1)
	}
	myAnimal := mal.Query(animalInfo[0])
	 switch animalInfo[1] {
	 case "eat":
		 fmt.Println("Eat: ",Eat(myAnimal))
	 case "move":
		 fmt.Println("Move:",Move(myAnimal))
	 case "speak":
		 fmt.Println("Speak:",Speak(myAnimal))
	 }
}

func(a *Animal) Query (animal string) *Animal{

	if animal == "cow" {
		cow := &Animal{
			Food:       "grass",
			Locomotion: "walk",
			Noise:      "moo",
		}
		return cow
	}

	if animal == "bird" {
		bird := &Animal{
			Food:       "worms",
			Locomotion: "fly",
			Noise:      "peep",
		}
		return bird
	}

	if animal == "snake" {
		snake := &Animal{
			Food:       "mice",
			Locomotion: "slither",
			Noise:      "hsss",
		}
		return snake
	}

	return nil
}
func Eat(myAnimal *Animal) string{
	return myAnimal.Food
}

func Move(myAnimal *Animal) string{
	return myAnimal.Locomotion
}

func Speak(myAnimal *Animal) string{
	return myAnimal.Noise
}