package main

import (
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/v4"
)

func main() {
e:
	echo.New()
	e.use(middleware.Logger())

}
