package main
import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type cow struct {
	name string
	food string
	locomotion string
	noise string
}

type bird struct {
	name string
	food string
	locomotion string
	noise string
}

type snake struct {
	name string
	food string
	locomotion string
	noise string
}

func (animal *cow) Eat() {
	fmt.Printf("%s\n",animal.food)	
}

func (animal *cow) Move() {
	fmt.Printf("%s\n",animal.locomotion)	
}

func (animal *cow) Speak() {
	fmt.Printf("%s\n",animal.noise)	
}

func (animal *bird) Eat() {
	fmt.Printf("%s\n",animal.food)	
}

func (animal *bird) Move() {
	fmt.Printf("%s\n",animal.locomotion)	
}

func (animal *bird) Speak() {
	fmt.Printf("%s\n",animal.noise)	
}

func (animal *snake) Eat() {
	fmt.Printf("%s\n",animal.food)	
}

func (animal *snake) Move() {
	fmt.Printf("%s\n",animal.locomotion)	
}

func (animal *snake) Speak() {
	fmt.Printf("%s\n",animal.noise)	
}


func main() {
	var cows []cow
	var birds []bird
	var snakes []snake

	var input_command string
	var input_name string
	var input_string2 string
	found:= false
	infinite := true

	fmt.Printf("x to Exit.\n")

	for infinite {
		fmt.Printf(">")
		fmt.Scanf("%s %s %s\n",&input_command,&input_name,&input_string2)
		switch input_command {
			case "newanimal":
				switch input_string2 {
					case "cow":
						tempCow := cow{input_name,"grass","walk","moo"}
						cows = append(cows,tempCow)
						fmt.Printf("Created it!\n")
					case "bird":
						tempBird := bird{input_name,"worms","fly","peep"}
						birds = append(birds,tempBird)
						fmt.Printf("Created it!\n")
					case "snake":
						tempSnake := snake{input_name,"mice","slither","hsss"}
						snakes = append(snakes,tempSnake)
						fmt.Printf("Created it!\n")
					default:
						fmt.Printf("Enter Valid animal type!\n")
				}
			case "query":
				for _, cow := range cows { //searching in cows
					if cow.name==input_name {
						switch input_string2 {
							case "eat":
								cow.Eat()
							case "move":
								cow.Move()
							case "speak":
								cow.Speak()
							default:
								fmt.Printf("Enter Valid action!\n")
						}
						found = true
					}	
				}

				for _, bird := range birds { //serching in birds
					if bird.name==input_name {
						switch input_string2 {
							case "eat":
								bird.Eat()
							case "move":
								bird.Move()
							case "speak":
								bird.Speak()
							default:
								fmt.Printf("Enter Valid action!\n")
						}
						found = true
					}	
				}

				for _, snake := range snakes { //searching in snakes 
					if snake.name==input_name {
						switch input_string2 {
							case "eat":
								snake.Eat()
							case "move":
								snake.Move()
							case "speak":
								snake.Speak()
							default:
								fmt.Printf("Enter Valid action!\n")
						}
						found = true
					}	
				}
				if !found {
					fmt.Printf("Animal Not Created with this name!\n")
				}
				found=false
			case "x":
				infinite=false
			default:
				fmt.Printf("Enter Valid command!\n")	
		}
	}
}