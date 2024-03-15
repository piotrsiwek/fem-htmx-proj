package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Templates


	}


func (t *Templates) Render(w echo.ResponseWriter, name string, data interface{}, c echo	Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate () *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

type Count struct {
	Count int

}
func main() {
e:
	echo.New()
	e.use(middleware.Logger())

	count = Count{Count: 0}
	e.Render= newTemplate()

	e.GET("/", func(c echo.Context) error {
		count.Count++
			return c.Render(200,"index.html",count)
})

    e.Logger.Fatal(e.Start(":8080"))

}