package main

import (
	"fmt"
	"purr"
)

func main() {
	fmt.Println("welcome to cattoDB")
	// fmt.Println(purr.Data)
	fmt.Println(purr.ReadOne("ajsonfile" + ".json"))
}
