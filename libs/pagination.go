package libs

import (
	"math"
	"strconv"
)

type Paginate struct {
	Total   int
	Page    int
	PerPage int
	Items   []interface{}
}

func (p *Paginate) Offset() int {
	return p.PerPage * (p.Page - 1)
}

func GetTotal(total int, per_page int) int {
	calculateTotal := strconv.Itoa(total / per_page)
	count, err := strconv.ParseFloat(calculateTotal, 64)
	if err != nil {
		return 0
	}
	ceiling := strconv.FormatFloat(math.Ceil(count), 'E', -1, 64)
	res, err := strconv.Atoi(ceiling)
	if err != nil {
		return 0
	}
	return res
}

func (p *Paginate) LastPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	if p.Page == total {
		return true
	}
	return false
}

func (p *Paginate) PrevPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	if p.Page <= total {
		return false
	}
	return true
}

func (p *Paginate) NextPage() bool {
	total := GetTotal(p.Total, p.PerPage)
	if p.Page <= total {
		return false
	}
	return true
}
