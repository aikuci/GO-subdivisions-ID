package model

import "github.com/lib/pq"

type CityResponse struct {
	BaseCollectionResponse[int]
	IDProvince  int           `json:"id_province"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	PostalCodes pq.Int64Array `json:"postal_codes"`
}

// ListCityByIDRequest defines a request structure for listing cities based on their ID.
type ListCityByIDRequest[T IdOrIds] struct {
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

// ListCityByIdRequest extends ListCityByIDRequest to support a different type for the province ID.
type ListCityByIdRequest[T IdOrIds, TProvince IdOrIds] struct {
	Include    []string  `json:"include" query:"include"`
	ID         T         `json:"-" params:"id" query:"id"`
	IDProvince TProvince `json:"-" params:"id_province" query:"id_province"`
}

// GetCityByIDRequest defines a request structure to retrieve a city based on their ID.
type GetCityByIDRequest[T IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}

// GetCityByIdRequest extends GetCityByIDRequest to support a different type for the province ID.
type GetCityByIdRequest[T IdOrIds, TProvince IdOrIds] struct {
	ID         T         `json:"-" params:"id" query:"id" validate:"required"`
	IDProvince TProvince `json:"-" params:"id_province" query:"id_province"`
	Include    []string  `json:"include" query:"include"`
}
