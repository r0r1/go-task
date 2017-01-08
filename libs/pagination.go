package libs

import "math"

type Paginate struct {
	Total   int
	Page    int
	PerPage int
	Limit   int
	Items   []interface{}
}

func (p *Paginate) Offset() int {
	return p.PerPage * (p.Page - 1)
}

func (p *Paginate) SetTotal(total int) {
	p.Total = total
}

func GetTotal(total int, per_page int) int {
	result := float64(total) / float64(per_page)
	return int(math.Ceil(result))
}

func (p *Paginate) LastPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	p.Limit = total
	return p.Page == total
}

func (p *Paginate) PrevPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	if p.Page == 0 || p.Page == 1 {
		return false
	}
	return p.Page <= total
}

func (p *Paginate) NextPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	return p.Page < total
}
