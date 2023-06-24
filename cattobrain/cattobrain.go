package cattobrain

import (
	"cattodb/controller"
	"fmt"
	"log"
	"net/http"
	"parser"

	"github.com/labstack/echo/v4"
)

func Meow() {
	// we will have to input the command
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
func MeowAPI(c echo.Context) error {
	command := c.Request().URL.Path
	commands, err := parser.GetCommand(command)
	if err != nil {
		log.Fatal(err)
	}
	result, err := controller.Controller(commands)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("returning value")
	return c.JSON(http.StatusOK, result)
}
