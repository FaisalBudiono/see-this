package main

import (
	"FaisalBudiono/see-this/app/db"
	"FaisalBudiono/see-this/app/httpren"
	"FaisalBudiono/see-this/app/strf"
	"FaisalBudiono/see-this/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	db.Init()

	e.Static("dist", "dist")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return httpren.Render(
			c,
			http.StatusOK,
			view.IndexPage(db.GetAllResources()),
		)
	})

	e.POST("/resources", func(c echo.Context) error {
		name := c.FormValue("name")
		slug := c.FormValue("slug")

		db.SaveResource(name, slug)

		return httpren.Render(
			c,
			http.StatusOK,
			view.ResourceList(db.GetAllResources()),
		)
	})

	e.POST("/verificator/resources/slug", func(c echo.Context) error {
		slug := strf.Slug(c.FormValue("name"))

		return httpren.Render(c, http.StatusOK, view.SlugForm(slug))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
