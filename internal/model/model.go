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

// Request
type ListRequest struct {
}
type GetByIDRequest[T IdOrIds] struct {
	ID T `json:"-" validate:"required"`
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
