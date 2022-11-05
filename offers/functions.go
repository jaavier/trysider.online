package offers

import "time"

type Offer struct {
	Id string `json:"id"`
	Description string `json:"description"`
	Store string `json:"store"`
	MaxClaims int `json:"maxClaims"`
	CreatedAt time.Time `json:"createdAt"`
}

func FindOffer(offerId string) Offer {
	return Offer{}
}

func CreateOffer(offer Offer) bool {
	return true
}

func DeleteOffer(offerId) bool {
	return true
}

