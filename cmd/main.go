package main

import (
	"FaisalBudiono/see-this/app/db"
	"FaisalBudiono/see-this/app/httpren"
	"FaisalBudiono/see-this/view"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	tempResources = make([]db.Resource, 0)
	id            = 2
)

func makeDummyResource() []db.Resource {
	return []db.Resource{
		{
			Id:   1,
			Name: "asd",
		},
		{
			Id:   2,
			Name: "kambing",
		},
	}
}

func main() {
	e := echo.New()

	tempResources = append(tempResources, makeDummyResource()...)

	e.Static("dist", "dist")

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return httpren.Render(c, http.StatusOK, view.IndexPage(tempResources))
	})

	e.POST("/resources", func(c echo.Context) error {
		name := c.FormValue("name")

		id += 1
		tempResources = append(tempResources, db.Resource{
			Id:   id,
			Name: name,
		})

		return httpren.Render(c, http.StatusOK, view.ResourceList(tempResources))
	})

	e.Logger.Fatal(e.Start(":8080"))
}
