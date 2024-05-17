package domain

type Paging struct {
	Page    int64 `json:"page" form:"page"`
	PerPage int64 `json:"per_page" form:"per_page"`
}
