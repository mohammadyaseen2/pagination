package pagination

import (
	"math"
)

func (p *Info) SetAllData(currentTotalElements int) {
	p.SetFirstPage()
	p.SetLastPage()
	p.SetCurrentPage(currentTotalElements)
	p.SetNextAndPreviousPage()
}

func (p *Info) SetFirstPage() {
	p.Resource.FirstPage = 0
}

func (p *Info) SetLastPage() {
	p.Resource.LastPage = p.Resource.TotalPages.(int) - 1
}

func (p *Info) SetCurrentPage(currentTotalElements int) {
	p.Resource.SizeItems = currentTotalElements
	p.Resource.CurrentPage = p.pagination.Page
}

func (p *Info) SetNextAndPreviousPage() {
	totalPages := p.Resource.TotalPages.(int) - 1
	if p.pagination.Page == totalPages {
		return
	} else if p.pagination.Page == 0 {
		p.SetNextPage()
	} else if p.pagination.Page < totalPages {
		p.SetPreviousPage()
		p.SetNextPage()
	} else if p.pagination.Page >= totalPages {
		p.SetPreviousPage()
	}
}

func (p *Info) SetPreviousPage() {
	p.Resource.PreviousPage = p.pagination.Page - 1
}

func (p *Info) SetNextPage() {
	p.Resource.NextPage = p.pagination.Page + 1
}

func getTotalPages(totalElements int, size int) int {
	if size == 0 {
		return 1
	}
	return int(math.Ceil(float64(totalElements) / float64(size)))
}
