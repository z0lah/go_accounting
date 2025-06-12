package pagination

import "github.com/gofiber/fiber/v2"

type PageQuery struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type Meta struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}
type PagedResponse[T any] struct {
	Data       []T  `json:"data"`
	Pagination Meta `json:"pagination"`
}

func (p *PageQuery) Normalize() {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}
	if p.Limit > 100 {
		p.Limit = 100
	}
}

// FromQuery parses query param into PageQuery
func FromQuery(c *fiber.Ctx) PageQuery {
	pq := PageQuery{}
	if err := c.QueryParser(&pq); err != nil {
		pq.Page = 1
		pq.Limit = 10
	}
	pq.Normalize()
	return pq
}

// BuildPagedResponse membungkus hasil paginasi
func BuildPagedResponse[T any](data []T, page int, limit int, total int64) PagedResponse[T] {
	return PagedResponse[T]{
		Data: data,
		Pagination: Meta{
			Page:  page,
			Limit: limit,
			Total: total,
		},
	}
}
