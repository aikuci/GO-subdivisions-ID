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

// ListDistrictByIDRequest defines a request structure for listing cities based on their ID.
type ListDistrictByIDRequest[T IdOrIds] struct {
	Include    []string `json:"include" query:"include"`
	ID         T        `json:"-" params:"id" query:"id"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
}

// ListDistrictByIdRequest extends ListDistrictByIDRequest to support a different types for province and city IDs.
type ListDistrictByIdRequest[T IdOrIds, TCity IdOrIds, TProvince IdOrIds] struct {
	Include    []string  `json:"include" query:"include"`
	ID         T         `json:"-" params:"id" query:"id"`
	IDCity     TCity     `json:"-" params:"id_city" query:"id_city"`
	IDProvince TProvince `json:"-" params:"id_province" query:"id_province"`
}

// GetDistrictByIDRequest defines a request structure to retrieve a city based on their ID.
type GetDistrictByIDRequest[T IdOrIds] struct {
	ID         T        `json:"-" params:"id" query:"id" validate:"required"`
	IDCity     T        `json:"-" params:"id_city" query:"id_city"`
	IDProvince T        `json:"-" params:"id_province" query:"id_province"`
	Include    []string `json:"include" query:"include"`
}

// GetDistrictByIdRequest extends GetDistrictByIDRequest to support a different types for province and city IDs.
type GetDistrictByIdRequest[T IdOrIds, TCity IdOrIds, TProvince IdOrIds] struct {
	ID         T         `json:"-" params:"id" query:"id" validate:"required"`
	IDCity     TCity     `json:"-" params:"id_city" query:"id_city"`
	IDProvince TProvince `json:"-" params:"id_province" query:"id_province"`
	Include    []string  `json:"include" query:"include"`
}
