package blacklist

import (
	"encoding/json"

	"github.com/jaavier/sider"
	"github.com/labstack/echo/v4"
)

type Body struct {
	Ip string `json:"ip"`
}

func Get(c echo.Context) error {
	bannedList, _ := sider.GetList("banned:ip")
	return c.JSON(200, bannedList)
}

func Post(c echo.Context) error {
	var body Body
	json.NewDecoder(c.Request().Body).Decode(&body)
	if len(body.Ip) == 0 {
		return c.String(400, "Value 'ip' is required")
	}
	var result bool = addIp(body.Ip)
	if result {
		return c.String(200, body.Ip+" banned!")
	} else {
		return c.String(400, body.Ip+" cannot be banned!")
	}
}

func Delete(c echo.Context) error {
	var body Body
	json.NewDecoder(c.Request().Body).Decode(&body)
	if len(body.Ip) == 0 {
		return c.String(400, "Value 'ip' is required")
	}
	if removeIp(body.Ip) {
		return c.String(200, "Unbanned!")
	}
	return c.String(400, "Error trying to remove ban")
}
