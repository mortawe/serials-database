package search

import "fmt"

type PageRequest struct {
	Page uint32 `json:"page" form:"page"`
	Size uint32 `json:"size" form:"size"`
}

func (r *PageRequest) ToSQL() string {
	return fmt.Sprintf(" OFFSET %d LIMIT %d ", (r.Page-1)*r.Size, r.Size)
}

type PageResponse struct {
	PageRequest
	Pages uint32 `json:"pages"`
}
