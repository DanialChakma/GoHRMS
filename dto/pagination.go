package dto

type PaginatedResponse[T any] struct {
	Data      []T   `json:"data"`
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"`
}
