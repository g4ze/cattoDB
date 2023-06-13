package controller

import (
	"errorhandle"
	"fmt"
	"log"
	"purr"
)

func Controller(commands []string) (purr.Employee, error) {
	if commands[0] == "read" { //readingone
		filep := "" + string(commands[1])
		fmt.Println(filep)
		if filep == "" {
			log.Fatal("error")
		}

		return purr.ReadOne(filep), nil
		// } else if command == "read-all" { //reading mutiple
		// 	files, err := ioutil.ReadDir(filep)
		// 	//got fie names just now
		// 	if err != nil {
		// 		fmt.Println("Error reading directory:", err)
		// 		return
		// 	} //iterating thru the file names
		// 	for _, file := range files {
		// 		if filepath.Ext(file.Name()) == ".json" {
		// 			path := filep + file.Name()
		// 			var fileslice purr.Employee

		// 		}

		// 	}
		// }
	}
	return purr.Employee{}, errorhandle.CustomError{Message: "invalid command meoww"}
}
