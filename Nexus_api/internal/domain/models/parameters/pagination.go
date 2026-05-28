package parameters

type PaginationQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

func NewPaginationQuery(page, size int) PaginationQuery {
	if page <= 0 {
		page = 1
	}

	if size < 12 {
		size = 12
	}

	return PaginationQuery{
		Page: page,
		Size: size,
	}
}