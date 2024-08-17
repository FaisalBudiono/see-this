package main

import (
	"FaisalBudiono/see-this/internal/app/adapter"
	"FaisalBudiono/see-this/internal/app/core/http/ren"
	"FaisalBudiono/see-this/internal/app/core/strf"
	"FaisalBudiono/see-this/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	db := adapter.NewDB()

	e.Static("dist", "dist")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		res, err := db.GetAllResources()
		if err != nil {
			return c.String(http.StatusInternalServerError, "something wrong")
		}

		return ren.Render(
			c,
			http.StatusOK,
			view.IndexPage(res),
		)
	})

	e.GET("/resources/:slug", func(c echo.Context) error {
		slug := c.Param("slug")

		res, err := db.GetAllResources()
		if err != nil {
			return c.String(http.StatusInternalServerError, "something wrong")
		}

		r, err := db.FindBySlug(slug)
		if err != nil {
			return ren.Render(
				c,
				http.StatusNotFound,
				view.ResourceNotFoundPage(res),
			)
		}

		return ren.Render(
			c,
			http.StatusOK,
			view.ResourcePage(res, r),
		)
	})

	e.POST("/resources", func(c echo.Context) error {
		name := c.FormValue("name")
		slug := c.FormValue("slug")

		_, err := db.SaveResource(name, slug)
		if err != nil {
			return c.String(http.StatusInternalServerError, "something wrong")
		}

		res, err := db.GetAllResources()
		if err != nil {
			return c.String(http.StatusInternalServerError, "something wrong")
		}

		return ren.Render(
			c,
			http.StatusOK,
			view.ResourceList(res),
		)
	})

	e.POST("/verificator/resources/slug", func(c echo.Context) error {
		slug := strf.Slug(c.FormValue("name"))

		return ren.Render(c, http.StatusOK, view.SlugForm(slug))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
