package utils

import (
	"fmt"
	"gin/entity/vo"
	"gin/model"
	"gin/types/relation"
	"math"
	"regexp"
)

type PageUtil struct {
	sql  string
	args []interface{}
	Page *relation.PageRequest
}

func New(sql string, args []interface{}, page *relation.PageRequest) *PageUtil {
	p := &PageUtil{}
	p.sql = sql
	p.args = args
	p.Page = page
	return p
}

func (p *PageUtil) StartPage() string {
	p.sql += fmt.Sprintf(" LIMIT %d OFFSET %d", p.Page.PageSize, (p.Page.Page-1)*p.Page.PageSize)
	return p.sql
}

func (p *PageUtil) EndPage(data any) *vo.PageVO {
	r1 := regexp.MustCompile(`(?i)SELECT .* FROM`)
	r2 := regexp.MustCompile(`LIMIT\s+\d+\s+OFFSET\s+\d+`)
	r3 := regexp.MustCompile(`(?i)\bOrder By\s+\S+(\s+ASC|\s+DESC)?\b`)
	countQuery := r3.ReplaceAllString(r2.ReplaceAllString(r1.ReplaceAllString(p.sql, "SELECT COUNT(*) FROM"), ""), "")

	var total int

	if p.args == nil {
		model.DB.Raw(countQuery).Scan(&total)
	} else {
		model.DB.Raw(countQuery, p.args).Scan(&total)
	}

	page := &vo.PageVO{}
	page.Page = p.Page.Page
	page.PageSize = p.Page.PageSize
	page.TotalPage = calculateTotalPages(total, page.PageSize)
	page.Data = data
	return page
}

func calculateTotalPages(totalRecords, recordsPerPage int) int {
	return int(math.Ceil(float64(totalRecords) / float64(recordsPerPage)))
}
