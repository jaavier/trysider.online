package offers

import (
	"encoding/json"
	"time"

	"github.com/labstack/echo/v4"
)

type Offer struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Store string `json:"store"`
	MaxClaims int `json:"maxClaims"`
	CreatedAt time.Time `json:"createdAt"`
}

func Get(c echo.Context) error {
	return c.String(200, "GET OFFER")
}

func Post(c echo.Context) error {
	var offer Offer
	json.NewDecoder(c.Request().Body).Decode(&offer)
	if CreateOffer(offer) {
		return c.String(200, "Offer created successfully")
	} else {
		return c.String(400, "Error creating offer")
	}
}

func Delete(c echo.Context) error {
	return c.String(200, "DELETE OFFER")
}
