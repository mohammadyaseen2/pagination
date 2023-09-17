package pagination

import "github.com/mohammadyaseen2/Pagination/model"

type Info struct {
	pagination *model.Pagination
	Resource   *model.Resource
}

func NewPaginationInfo(totalElements, page, size int) *Info {
	pagination := newPagination(page, size)
	return &Info{
		pagination: pagination,
		Resource: &model.Resource{
			TotalItems: totalElements,
			TotalPages: getTotalPages(totalElements, pagination.Size),
		},
	}
}

func newPagination(page int, size int) *model.Pagination {
	return &model.Pagination{
		Page: page,
		Size: size,
	}
}
