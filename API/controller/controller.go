package controller

import (
	"cattobrain"

	"github.com/labstack/echo"
)

// func ReadOne(w http.ResponseWriter, r *http.Request) {

// 	// controller.Controller( []string)
// 	// Process the identifier
// 	// You can perform any desired logic with the identifier
// 	command := (r.URL.Path[1:])
// 	result := cattobrain.MeowAPI(command)
// 	w.Header().Set("Content-Type", "application/json")
// 	response, _ := json.Marshal(result)
// 	w.Write(response)
// }
func ReadOne(context echo.Context) error {
	command := context.Param("id")
	result := cattobrain.MeowAPI(command)
}
