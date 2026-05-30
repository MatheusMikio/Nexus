package parameters

type PaginationQuery struct {
	Page int `form:"page"`
	Size int `form:"size"`
}

const (
	defaultPage = 1
	defaultSize = 12
	maxSize     = 20
)

func NewPaginationQuery(page, size int) PaginationQuery {
	if page <= 0 {
		page = defaultPage
	}

	if size <= 0 {
		size = defaultSize
	}

	if size > maxSize {
		size = maxSize
	}

	return PaginationQuery{
		Page: page,
		Size: size,
	}
}
