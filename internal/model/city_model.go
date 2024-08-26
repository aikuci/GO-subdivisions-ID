package model

import "github.com/lib/pq"

type CityResponse struct {
	BaseCollectionResponse[int]
	IDProvince  int                `json:"id_province"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	PostalCodes pq.Int64Array      `json:"postal_codes"`
	Districts   []DistrictResponse `json:"districts,omitempty"`
}

type ListCityByIDRequest[T IdPlural] struct {
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

type GetCityByIDRequest[T IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}
