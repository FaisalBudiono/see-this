package main

import (
	"FaisalBudiono/see-this/view"
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("dist", "dist")

	e.GET("/", func(c echo.Context) error {
		return Render(c, http.StatusOK, view.Kuda("asd"))
	})
	e.Logger.Fatal(e.Start(":8000"))
}

func Render(ctx echo.Context, statusCode int, t templ.Component) error {
	buf := templ.GetBuffer()
	defer templ.ReleaseBuffer(buf)

	if err := t.Render(ctx.Request().Context(), buf); err != nil {
		return err
	}

	return ctx.HTML(statusCode, buf.String())
}
