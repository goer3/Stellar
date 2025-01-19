package dto

// 用户列表筛选条件
type SystemUserFilterRequest struct {
	JobNumber         *string `json:"jobNumber" form:"jobNumber"`
	Username          *string `json:"username" form:"username"`
	CnName            *string `json:"cnName" form:"cnName"`
	EnName            *string `json:"enName" form:"enName"`
	Email             *string `json:"email" form:"email"`
	Phone             *string `json:"phone" form:"phone"`
	Status            *uint   `json:"status" form:"status"`
	Gender            *uint   `json:"gender" form:"gender"`
	SystemRole        *uint   `json:"systemRole" form:"systemRole"`
	SystemDepartment  *uint   `json:"systemDepartment" form:"systemDepartment"`
	SystemJobPosition *uint   `json:"systemJobPosition" form:"systemJobPosition"`
	PageNumber        *uint   `json:"pageNumber" form:"pageNumber"`
	PageSize          *uint   `json:"pageSize" form:"pageSize"`
	Paginable         *bool   `json:"paginable" form:"paginable"`
}
