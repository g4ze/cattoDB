package cattobrain

import (
	"fmt"
)

func ReadCommand() string {
	var command string

	fmt.Print("( • ω • )=>")
	fmt.Scanln(&command)
	return command
}
