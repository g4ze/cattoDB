package parser

import "errorhandle"

func GetCommand(command string) (string, error) {
	//this function will be parsing and breaking down the commands.
	if command == "" {
		return "", errorhandle.CustomError{Message: "no command found"}
	}
	if command[0:6] != "catto-" {
		return "", errorhandle.CustomError{Message: "try using \"catto\" command and \"catto-help\" for more"}
	}
	command = command[6:]
	return command, nil
}
