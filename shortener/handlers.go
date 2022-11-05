package shortener

import "github.com/labstack/echo/v4"

func Get(c echo.Context) error {
	return c.String(200, "Redirecting")
}

func Post(c echo.Context) error {
	return c.String(200, "Create short URL")
}

func Delete(c echo.Context) error {
	return c.String(200, "Delete short URL")
}
