package model

import "github.com/lib/pq"

type CityResponse struct {
	BaseCollectionResponse[int]
	IDProvince  int           `json:"id_province"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	PostalCodes pq.Int64Array `json:"postal_codes"`
}

// GetCityByIDRequest defines a request structure to retrieve a city by its ID.
// The generic type T allows for different ID types (single or multiple).
// Both city ID and province ID are required parameters for this request.
type GetCityByIDRequest[T IdOrIds] struct {
	ID         T `json:"-" validate:"required"`
	IDProvince T `json:"-" params:"id_province" query:"id_province" validate:"required"`
}

// GetCityByIdRequest extends GetCityByIDRequest to support a different type for the province ID.
// This structure allows for more flexibility with different types for city and province IDs.
// Both city ID and province ID are required parameters.
type GetCityByIdRequest[T IdOrIds, TProvince IdOrIds] struct {
	Id         T         `json:"-" validate:"required"`
	IdProvince TProvince `json:"-" params:"id_province" query:"id_province" validate:"required"`
}

// GetCityByIDProvinceRequest defines a request structure to retrieve cities by province ID only.
// This request does not require a city ID, focusing on filtering cities based on the province ID alone.
type GetCityByIDProvinceRequest[T IdOrIds] struct {
	IDProvince T `json:"-" params:"id_province" query:"id_province"`
}
