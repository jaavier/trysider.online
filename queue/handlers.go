package queue

import "github.com/labstack/echo/v4"

func Get(c echo.Context) error {
	return c.String(200, "GET QUEUE")
}

func Post(c echo.Context) error {
	return c.String(200, "GET QUEUE")
}

func Delete(c echo.Context) error {
	return c.String(200, "GET QUEUE")
}
