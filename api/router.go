package main

import (
	"api/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setRouter() (err error) {
	db, err := newDB()
	if err != nil {
		return
	}

	e := echo.New()

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
	)

	e.GET("/photos", func(c echo.Context) error { return controller.GetPhotos(c, db) })
	e.GET("/collections", func(c echo.Context) error { return controller.GetCollections(c, db) })

	e.Logger.Fatal(e.Start(":9000"))

	return
}
