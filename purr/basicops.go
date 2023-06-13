package purr

import (
	"encoding/json"
	"fmt"
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
	var datatype Employee
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	// _ = json.NewDecoder(file).Decode(&datatype)
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &datatype)
	return datatype
}
