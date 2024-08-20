package model

import "github.com/lib/pq"

type ProvinceResponse struct {
	ID          int           `json:"id"`
	Code        string        `json:"code"`
	Name        string        `json:"name"`
	PostalCodes pq.Int64Array `json:"postal_codes"`
}

type ListProvinceRequest struct {
}

type GetProvinceRequest struct {
	ID int `json:"-" validate:"required"`
}
