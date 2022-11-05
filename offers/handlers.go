package offers

import "github.com/labstack/echo/v4"

func Get(c echo.Context) error {
	return c.String(200, "GET OFFER")
}

func Post(c echo.Context) error {
	return c.String(200, "CREATE OFFER")

}

func Delete(c echo.Context) error {
	return c.String(200, "DELETE OFFER")
}
