package model

import (
	appmodel "github.com/aikuci/go-subdivisions-id/pkg/model"

	"github.com/lib/pq"
)

type VillageResponse struct {
	appmodel.BaseCollectionResponse[int]
	IDDistrict  int              `json:"id_district"`
	IDCity      int              `json:"id_city"`
	IDProvince  int              `json:"id_province"`
	Code        string           `json:"code"`
	Name        string           `json:"name"`
	PostalCodes pq.Int64Array    `json:"postal_codes"`
	District    DistrictResponse `json:"district,omitempty"`
	City        CityResponse     `json:"city,omitempty"`
	Province    ProvinceResponse `json:"province,omitempty"`
}

type ListVillageByIDRequest[T appmodel.IdPlural] struct {
	appmodel.PageRequest
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDDistrict T        `json:"-" params:"id_district" query:"id_district"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

type GetVillageByIDRequest[T appmodel.IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDDistrict T        `json:"-" params:"id_district" query:"id_district"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}
