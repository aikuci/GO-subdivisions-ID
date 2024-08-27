package model

import (
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"github.com/lib/pq"
)

type CityResponse struct {
	appmodel.BaseCollectionResponse[int]
	IDProvince  int                `json:"id_province"`
	Code        string             `json:"code"`
	Name        string             `json:"name"`
	PostalCodes pq.Int64Array      `json:"postal_codes"`
	Province    ProvinceResponse   `json:"province,omitempty"`
	Districts   []DistrictResponse `json:"districts,omitempty"`
}

type ListCityByIDRequest[T appmodel.IdPlural] struct {
	appmodel.PageRequest
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

type GetCityByIDRequest[T appmodel.IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}
