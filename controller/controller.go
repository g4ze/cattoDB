package controller

import (
	"errorhandle"
	"purr"
)

func controller(command string, filep string) (purr.Employee, error) {
	if command == "read" { //readingone
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
