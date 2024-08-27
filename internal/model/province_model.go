package model

import (
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"github.com/lib/pq"
)

type ProvinceResponse struct {
	appmodel.BaseCollectionResponse[int]
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	PostalCodes pq.Int64Array      `json:"postal_codes"`
	Cities      []CityResponse     `json:"cities,omitempty"`
	Districts   []DistrictResponse `json:"districts,omitempty"`
}
