package controller

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"purr"
)

func controller(command string, filep string) (a purr.Employee) {
	if command == "read" { //readingone
		return purr.ReadOne(filep)
	} else if command == "read-all" { //reading mutiple
		files, err := ioutil.ReadDir(filep)
		//got fie names just now
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		} //iterating thru the file names
		for _, file := range files {
			if filepath.Ext(file.Name()) == ".json" {
				path := filep + file.Name()
				var fileslice purr.Employee

			}

		}
	}
}
