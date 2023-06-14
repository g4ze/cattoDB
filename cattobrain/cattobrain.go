package cattobrain

import (
	"cattodb/controller"
	"fmt"
	"log"
	"parser"
	"purr"
)

func Meow() {
	//we will have to input the command
	command := ReadCommand()

	commands, err := parser.GetCommand(command)
	if err != nil {
		log.Fatal(err)
	}
	result, err := controller.Controller(commands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result)
}
func MeowAPI(command string) purr.Employee {
	commands, err := parser.GetCommand(command)
	if err != nil {
		log.Fatal(err)
	}
	result, err := controller.Controller(commands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("returning value")
	return result
}
