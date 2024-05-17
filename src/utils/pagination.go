// Package pagination provides support for pagination requests and responses.
package utils

import (
	"fmt"
	"strings"
)

var (
	// DefaultPageSize specifies the default page size
	DefaultPageSize int64 = 20
	// MaxPageSize specifies the maximum page size
	MaxPageSize int64 = 100
	// PageVar specifies the query parameter name for page number
	PageVar = "page"
	// PageSizeVar specifies the query parameter name for page size
	PageSizeVar = "per_page"
)

// Pagination represents a paginated list of data items.
type Pagination struct {
	Page       int64       `json:"page" form:"page"`
	PerPage    int64       `json:"per_page" form:"per_page"`
	PageCount  int64       `json:"page_count" form:"_"`
	TotalCount int64       `json:"total_count" form:"_"`
	Data       interface{} `json:"data" form:"_"`
}

// New creates a new Pages instance.
// The page parameter is 1-based and refers to the current page index/number.
// The perPage parameter refers to the number of items on each page.
func NewPagination(page, perPage int64) *Pagination {
	if perPage <= 0 {
		perPage = DefaultPageSize
	}
	if perPage > MaxPageSize {
		perPage = MaxPageSize
	}
	if page < 1 {
		page = 1
	}

	return &Pagination{
		Page:    page,
		PerPage: perPage,
	}
}

func (p *Pagination) SetTotal(total int64) {
	var pageCount int64 = -1
	if total >= 0 {
		pageCount = (total + p.PerPage - 1) / p.PerPage
		if p.Page > pageCount {
			p.Page = pageCount
		}
	}
}

// Offset returns the OFFSET value that can be used in a SQL statement.
func (p *Pagination) Offset() int64 {
	return (p.Page - 1) * p.PerPage
}

// Limit returns the LIMIT value that can be used in a SQL statement.
func (p *Pagination) Limit() int64 {
	return p.PerPage
}

// BuildLinkHeader returns an HTTP header containing the links about the pagination.
func (p *Pagination) BuildLinkHeader(baseURL string, defaultPerPage int64) string {
	links := p.BuildLinks(baseURL, defaultPerPage)
	header := ""
	if links[0] != "" {
		header += fmt.Sprintf("<%v>; rel=\"first\", ", links[0])
		header += fmt.Sprintf("<%v>; rel=\"prev\"", links[1])
	}
	if links[2] != "" {
		if header != "" {
			header += ", "
		}
		header += fmt.Sprintf("<%v>; rel=\"next\"", links[2])
		if links[3] != "" {
			header += fmt.Sprintf(", <%v>; rel=\"last\"", links[3])
		}
	}
	return header
}

// BuildLinks returns the first, prev, next, and last links corresponding to the pagination.
// A link could be an empty string if it is not needed.
// For example, if the pagination is at the first page, then both first and prev links
// will be empty.
func (p *Pagination) BuildLinks(baseURL string, defaultPerPage int64) [4]string {
	var links [4]string
	pageCount := p.PageCount
	page := p.Page
	if pageCount >= 0 && page > pageCount {
		page = pageCount
	}
	if strings.Contains(baseURL, "?") {
		baseURL += "&"
	} else {
		baseURL += "?"
	}
	if page > 1 {
		links[0] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, 1)
		links[1] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page-1)
	}
	if pageCount >= 0 && page < pageCount {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
		links[3] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, pageCount)
	} else if pageCount < 0 {
		links[2] = fmt.Sprintf("%v%v=%v", baseURL, PageVar, page+1)
	}
	if perPage := p.PerPage; perPage != defaultPerPage {
		for i := 0; i < 4; i++ {
			if links[i] != "" {
				links[i] += fmt.Sprintf("&%v=%v", PageSizeVar, perPage)
			}
		}
	}

	return links
}
