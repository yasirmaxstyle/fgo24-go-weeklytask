package utils

const ItemsPerPage = 3

type Pagination struct {
	CurrentPage  int
	ItemsPerPage int
	TotalItems   int
}

func NewPagination(totalItems int) *Pagination {
	return &Pagination{
		CurrentPage:  0,
		ItemsPerPage: ItemsPerPage,
		TotalItems:   totalItems,
	}
}

func (p *Pagination) GetCurrentPageItems(items []MenuItem) []MenuItem {
	startIdx := p.CurrentPage * p.ItemsPerPage
	endIdx := min(startIdx+p.ItemsPerPage, p.TotalItems)

	if startIdx >= p.TotalItems {
		return []MenuItem{}
	}

	return items[startIdx:endIdx]
}

func (p *Pagination) NextPage() bool {
	totalPages := (p.TotalItems + p.ItemsPerPage - 1) / p.ItemsPerPage
	if p.CurrentPage < totalPages-1 {
		p.CurrentPage++
		return true
	}
	return false
}

func (p *Pagination) PreviousPage() bool {
	if p.CurrentPage > 0 {
		p.CurrentPage--
		return true
	}
	return false
}

func (p *Pagination) GetTotalPages() int {
	return (p.TotalItems + p.ItemsPerPage - 1) / p.ItemsPerPage
}

func (p *Pagination) GetStartIndex() int {
	return p.CurrentPage * p.ItemsPerPage
}

func (p *Pagination) GetEndIndex() int {
	endIdx := min(p.GetStartIndex()+p.ItemsPerPage, p.TotalItems)
	return endIdx
}
