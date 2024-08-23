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
type ListRequest struct {
}
type GetByIDRequest[T IdOrIds] struct {
	// The GET method does not include a request body.
	// Therefore, we skip the bodyParser method and omit the json struct tag for this field.
	ID T `json:"-" params:"id" query:"id" validate:"required"`
}

// Response
type BaseCollectionResponse[T IdSingular] struct {
	ID T `json:"id"`
}
type WebResponse[T any] struct {
	Data   T             `json:"data"`
	Paging *PageMetadata `json:"paging,omitempty"`
	Errors string        `json:"errors,omitempty"`
}
type PageResponse[T any] struct {
	Data         []T          `json:"data,omitempty"`
	PageMetadata PageMetadata `json:"paging,omitempty"`
}
type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
