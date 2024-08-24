package model

import "github.com/lib/pq"

type ProvinceResponse struct {
	BaseCollectionResponse[int]
	Code        string         `json:"code"`
	Name        string         `json:"name"`
	PostalCodes pq.Int64Array  `json:"postal_codes"`
	Cities      []CityResponse `json:"cities,omitempty"`
}
