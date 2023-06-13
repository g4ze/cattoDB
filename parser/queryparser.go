package parser

import (
	"errorhandle"
	"fmt"
	"strings"
)

func GetCommand(command string) ([]string, error) {
	//this function will be parsing and breaking down the commands.
	if command == "" {
		return nil, errorhandle.CustomError{Message: "no command found"}
	}
	if command[0:6] != "catto-" {
		return nil, errorhandle.CustomError{Message: "try using \"catto\" command and \"catto-help\" for more"}
	}
	command = command[6:]

	commands := strings.Split(command, "-")
	var list *[]string = &commands
	fmt.Println("List of commands is ", *list)
	return *list, nil
}
