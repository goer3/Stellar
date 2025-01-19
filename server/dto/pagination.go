package dto

import "stellar/common"

// 分页信息
type PaginationConfig struct {
	PageNumber uint  `json:"pageNumber"` // 页码
	PageSize   uint  `json:"pageSize"`   // 每页数据量
	Total      int64 `json:"total"`      // 数据量
	Paginable  bool  `json:"paginable"`  // 是否需要分页，默认 false，则不需要分页
}

// 分页参数处理，获取分页查询的 limit 和 offset
func (p *PaginationConfig) GetPaginationLimitAndOffset() (limit, offset int) {
	// 如果不需要分页，则直接返回
	if !p.Paginable {
		return 0, 0
	}

	// 默认每页显示的数量是用户传递的数量，但是如果每页显示数量 < 1，或者超出最大的限制，则使用默认分页数
	pageSize := p.PageSize
	if pageSize < 1 || pageSize > common.SYSTEM_PAGE_MAX_SIZE {
		pageSize = common.SYSTEM_PAGE_DEFAULT_SIZE
	}

	// 默认页码是用户传递的页码，但是如果传递的页面 -1，则返回第一页，至于传递很大的页面，不用管，最终返回没数据就行
	pageNumber := p.PageNumber
	if pageNumber < 1 {
		pageNumber = 1
	}

	// 回填数据
	p.PageSize = pageSize
	p.PageNumber = pageNumber

	// 返回 limit 和 offset，使用 int 类型是为了和数据库查询的 limit 和 offset 保持一致
	return int(pageSize), int((pageNumber - 1) * pageSize)
}

// 数据分页返回格式
type PaginationResponse struct {
	Pagination PaginationConfig `json:"pagination"`
	List       interface{}      `json:"list"`
}
