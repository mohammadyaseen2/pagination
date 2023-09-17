package pagination

import "github.com/mohammadyaseen2/Pagination/model"

type Info struct {
	pagination *model.Pagination
	Resource   *model.Resource
}

func NewPaginationInfo(totalElements int, pagination *model.Pagination) *Info {
	return &Info{
		pagination: pagination,
		Resource: &model.Resource{
			TotalItems: totalElements,
			TotalPages: getTotalPages(totalElements, pagination.Size),
		},
	}
}

func NewPagination(page int, size int) *model.Pagination {
	return &model.Pagination{
		Page: page,
		Size: size,
	}
}
