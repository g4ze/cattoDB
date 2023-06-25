package router

import (
	"cattobrain"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route() {
	// router := mux.NewRouter()
	// fmt.Println("Route was hit")
	// router.HandleFunc("/catto-read-{id}", controller.ReadOne).Methods("GET")
	// router.HandleFunc("/catto-read-all", controller.ReadOne).Methods("GET")
	// router.HandleFunc("/catto-read-{ids}", controller.ReadOne).Methods("GET")

	// // Creating a new server
	// server := &http.Server{
	// 	Addr:         "localhost:" + port,
	// 	ReadTimeout:  3 * time.Second,
	// 	WriteTimeout: 6 * time.Second,

	// 	IdleTimeout:       15 * time.Second,
	// 	ReadHeaderTimeout: 3 * time.Second,
	// 	Handler:           router,
	// }

	// // Start the server
	// log.Fatal(server.ListenAndServe())
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "meowww :3! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/:id", func(c echo.Context) error {
		return cattobrain.MeowAPI(c)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = port
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
