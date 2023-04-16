package main

import (
	"log"
	"net/http"

	controller "github.com/Rizqirf/go_echo/controllers"
	seed "github.com/Rizqirf/go_echo/seeder"
	"github.com/Rizqirf/go_echo/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	storage.DbConnect()
	seeder()
	// Middleware
	// e.Use(middleware.Logger())
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", healthCheck)
	e.GET("/property-promotion", controller.GetAll)
	e.GET("/property-promotion/:id", controller.GetById)
	e.POST("/property-promotion",controller.Create)
	e.PUT("/property-promotion/:id",controller.Update)
	e.DELETE("/property-promotion/:id",controller.Delete)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

// Handler
func healthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Server Running!")
}

//seeder
func seeder(){
	db:=storage.GetDBInstance()
	for _, seed := range seed.All() {
		if err := seed.Run(db); err != nil {
			log.Fatalf("Running seed '%s', failed with error: %s", seed.Name, err)
		}
	}
}