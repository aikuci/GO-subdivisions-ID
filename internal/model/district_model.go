package model

import (
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"github.com/lib/pq"
)

type DistrictResponse struct {
	appmodel.BaseCollectionResponse[int]
	IDCity      int              `json:"id_city"`
	IDProvince  int              `json:"id_province"`
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	PostalCodes pq.Int64Array    `json:"postal_codes"`
	City        CityResponse     `json:"city,omitempty"`
	Province    ProvinceResponse `json:"province,omitempty"`
}

type ListDistrictByIDRequest[T appmodel.IdPlural] struct {
	appmodel.PageRequest
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

type GetDistrictByIDRequest[T appmodel.IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}
