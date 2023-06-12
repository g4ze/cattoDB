package purr

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func WriteOne(filepath string, data Employee) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(filepath, file, 0644)

}
func DeleteOne(filepath string) {
	e := os.Remove(filepath)
	if e != nil {
		log.Fatal(e)
	}
}
func ReadOne(filepath string) (a Employee) {
	file, _ := ioutil.ReadFile(filepath)
	var datatype Employee

	_ = json.Unmarshal([]byte(file), &datatype)
	return datatype
}
