package model

import "github.com/lib/pq"

type DistrictResponse struct {
	BaseCollectionResponse[int]
	IDCity      int           `json:"id_city"`
	IDProvince  int           `json:"id_province"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	PostalCodes pq.Int64Array `json:"postal_codes"`
}

type ListDistrictByIDRequest[T IdPlural] struct {
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

type GetDistrictByIDRequest[T IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}
