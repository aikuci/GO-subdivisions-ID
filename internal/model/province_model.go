package model

type ProvinceResponse struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	PostalCodes []int  `json:"postal_codes"`
}

type ListProvinceRequest struct {
}

type GetProvinceRequest struct {
	ID int `json:"-" validate:"required"`
}
