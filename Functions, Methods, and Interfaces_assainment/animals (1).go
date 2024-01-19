package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var valsS []string
var dict = make(map[string]Animal)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Printf("grass\n")
}

func (c Cow) Move() {
	fmt.Printf("walk\n")
}

func (c Cow) Speak() {
	fmt.Printf("moo\n")
}

type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Printf("worms\n")
}

func (b Bird) Move() {
	fmt.Printf("fly\n")
}

func (b Bird) Speak() {
	fmt.Printf("peep\n")
}

type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Printf("mice\n")
}

func (s Snake) Move() {
	fmt.Printf("slither\n")
}

func (s Snake) Speak() {
	fmt.Printf("hsss\n")
}

func answerQuery(name, info string) {
	animal := dict[name]
	if animal != nil {
		if info != "eat" && info != "move" && info != "speak" {
			fmt.Printf("Remember you can ask for 'eat', 'move' and 'speak'\n")
		} else {
			switch info {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		}
	} else {
		fmt.Println("This animal doesnt' exist yet.")
	}
}

func createAnimal(name, animal string) {
	if animal != "cow" && animal != "bird" && animal != "snake" {
		fmt.Printf("Remember you can create 'cow', 'bird' or 'snake'\n")
	} else {
		switch animal {
		case "cow":
			cow := Cow{name: name}
			dict[name] = cow
		case "bird":
			bird := Bird{name: name}
			dict[name] = bird
		case "snake":
			snake := Snake{name: name}
			dict[name] = snake
		}
		fmt.Printf("Created it!\n")
	}

}

func createOrQuery() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("> ")
	if scanner.Scan() {
		valsS = strings.Split(scanner.Text(), " ")
		if len(valsS) == 3 {
			if valsS[0] == "newanimal" {
				createAnimal(valsS[1], valsS[2])
			} else if valsS[0] == "query" {
				answerQuery(valsS[1], valsS[2])
			} else {

			}
		} else {
			createOrQuery()
		}

	}
}

func main() {
	fmt.Println("Example for creating an animal: 'newanimal lulu cow' ")
	fmt.Println("Example for getting info of an animal: 'query lulu eat' ")
	for {
		createOrQuery()
	}
}
