package model

import "github.com/lib/pq"

type CityResponse struct {
	BaseCollectionResponse
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	PostalCodes pq.Int64Array `json:"postal_codes"`
}
