package main

import (
	"FaisalBudiono/see-this/internal/db"
	"FaisalBudiono/see-this/internal/http/ren"
	"FaisalBudiono/see-this/internal/strf"
	"FaisalBudiono/see-this/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Static("dist", "dist")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return ren.Render(
			c,
			http.StatusOK,
			view.IndexPage(db.GetAllResources()),
		)
	})

	e.GET("/resources/:slug", func(c echo.Context) error {
		slug := c.Param("slug")

		r, err := db.FindBySlug(slug)
		if err != nil {
			return ren.Render(
				c,
				http.StatusNotFound,
				view.ResourceNotFoundPage(db.GetAllResources()),
			)
		}

		return ren.Render(
			c,
			http.StatusOK,
			view.ResourcePage(db.GetAllResources(), r),
		)
	})

	e.POST("/resources", func(c echo.Context) error {
		name := c.FormValue("name")
		slug := c.FormValue("slug")

		db.SaveResource(name, slug)

		return ren.Render(
			c,
			http.StatusOK,
			view.ResourceList(db.GetAllResources()),
		)
	})

	e.POST("/verificator/resources/slug", func(c echo.Context) error {
		slug := strf.Slug(c.FormValue("name"))

		return ren.Render(c, http.StatusOK, view.SlugForm(slug))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
