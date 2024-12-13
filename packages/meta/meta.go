package meta

import (
	"os"
	"strconv"
)

type Meta struct {
	TotalCount int `json:"total_records"`
	PageCount  int `json:"total_pages"`
	Page       int `json:"current_page"`
	PerPage    int `json:"per_page"`
}

func New(page, perPage, total int) (*Meta, error) {
	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(os.Getenv("PAGINATOR_LIMIT_DEFAULT"))
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		TotalCount: total,
		PageCount:  pageCount,
	}, nil
}

func (m *Meta) Offset() int {
	return (m.Page - 1) * m.PerPage
}

func (m *Meta) Limit() int {
	return m.PerPage
}
