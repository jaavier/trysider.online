package offers

import (
	"strconv"

	"github.com/jaavier/sider"
)

func FindOffer(offerId string) Offer {
	_, err := sider.Get("offer:id:" + offerId)
	if err != nil {
		return Offer{}	
	}
	description, _ := sider.Get("offer:description:" + offerId)
	store, _ := sider.Get("offer:store:" + offerId)
	maxClaims, _ := sider.Get("offer:maxClaims:" + offerId)
	maxClaimsInt, _ := strconv.Atoi(maxClaims)
	// createdAt, _ := sider.Get("offer:createdAt:" + offerId)
	return Offer{
		Description: description,
		Store: store,
		MaxClaims: maxClaimsInt,
	}
}

func CreateOffer(offer Offer) bool {
	if offer.MaxClaims > 0 {
		return true
	} else {
		return false
	}
}

func DeleteOffer(offerId string) bool {
	if len(offerId) > 0 {
		return true
	} else {
		return false
	}
}
