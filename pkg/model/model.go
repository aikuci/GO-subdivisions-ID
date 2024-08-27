package model

import "github.com/google/uuid"

// ID
type IdSingular interface {
	int | string | uuid.UUID
}
type IdPlural interface {
	[]int | []string | []uuid.UUID
}
type IdOrIds interface {
	IdSingular | IdPlural
}
type IdSingularRequest[T IdSingular] struct {
	ID T `json:"id" params:"id" query:"id" validate:"required"`
}
type IdPluralRequest[T IdPlural] struct {
	ID T `json:"id" params:"id" query:"id" validate:"required"`
}

// Request
type PageRequest struct {
	Page int `json:"page" query:"page"`
	Size int `json:"size" query:"size"`
}
type ListRequest struct {
	PageRequest
	Include []string `json:"include" query:"include"`
}
type GetByIDRequest[T IdOrIds] struct {
	Include []string `json:"include" query:"include"`
	ID      T        `json:"-" params:"id" query:"id" validate:"required"`
}

// Response
type BaseCollectionResponse[T IdSingular] struct {
	ID T `json:"id"`
}
type WebResponse[T any] struct {
	Data   T      `json:"data"`
	Meta   *Meta  `json:"meta,omitempty"`
	Errors string `json:"errors,omitempty"`
}
type Meta struct {
	Page *PageMetadata `json:"page,omitempty"`
}
type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"page,omitempty"`
}
type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
