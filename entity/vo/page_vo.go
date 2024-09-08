package vo

type PageVO struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	Data      any `json:"data"`
}

func (p *PageVO) Build() {

}
