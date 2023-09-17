package model

type Pagination struct {
	Page int
	Size int
}

type Resource struct {
	FirstPage    any `json:"firstPage,omitempty"`
	NextPage     any `json:"nextPage,omitempty"`
	CurrentPage  any `json:"currentPage,omitempty"`
	PreviousPage any `json:"previousPage,omitempty"`
	LastPage     any `json:"lastPage,omitempty"`
	SizeItems    any `json:"sizeItems,omitempty"`
	TotalPages   any `json:"totalPages,omitempty"`
	TotalItems   any `json:"totalItems,omitempty"`
}

type PaginationI interface {
	SetAllData(int)
	SetFirstPage()
	SetLastPage()
	SetCurrentPage(int)
	SetNextAndPreviousPage()
	SetPreviousPage()
	SetNextPage()
}
