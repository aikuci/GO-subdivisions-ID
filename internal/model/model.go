package model

type BaseCollectionResponse struct {
	ID int `json:"id"`
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

type IdOrIds interface {
	int | []int
}

type ListRequest struct {
}

type GetByIDRequest[T IdOrIds] struct {
	ID T `json:"-" validate:"required"`
}
