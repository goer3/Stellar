package model

import "github.com/dromara/carbon/v2"

// 用户模型
type SystemUser struct {
	BaseModel
	JobNumber            string              `gorm:"uniqueIndex:idx_jobNumber;comment:工号" json:"jobNumber"`
	Username             string              `gorm:"uniqueIndex:idx_username;comment:用户名" json:"username"`
	CNName               string              `gorm:"comment:中文名" json:"cnName"`
	ENName               string              `gorm:"comment:英文名" json:"enName"`
	Password             string              `gorm:"not null;comment:密码" json:"password"`
	Email                string              `gorm:"uniqueIndex:idx_email;comment:邮箱" json:"email"`
	Phone                string              `gorm:"uniqueIndex:idx_phone;comment:手机号" json:"phone"`
	HidePhone            *uint               `gorm:"default:0;comment:是否隐藏手机号，0：否，1：是" json:"hidePhone"`
	Gender               *uint               `gorm:"default:0;comment:性别，0：未知，1：男，2：女" json:"gender"`
	Avatar               string              `gorm:"not null;comment:头像" json:"avatar"`
	LastLoginIP          string              `gorm:"comment:最后登录IP" json:"lastLoginIP"`
	LastLoginAt          carbon.Carbon       `gorm:"comment:最后登录时间" json:"lastLoginAt"`
	LastChangePasswordAt carbon.Carbon       `gorm:"comment:最后修改密码时间" json:"lastChangePasswordAt"`
	Status               *uint               `gorm:"default:1;comment:状态，1：正常，0：禁用" json:"status"`
	Creator              string              `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	SystemRoleId         uint                `gorm:"comment:角色ID" json:"systemRoleId"`
	SystemRole           *SystemRole         `gorm:"foreignKey:SystemRoleId;" json:"systemRole,omitempty"` // 指针类型解决 omitempty 为空问题
	SystemJobPositions   []SystemJobPosition `gorm:"many2many:system_user_job_position_relation;" json:"systemJobPositions,omitempty"`
	SystemDepartments    []SystemDepartment  `gorm:"many2many:system_user_department_relation;" json:"systemDepartments,omitempty"`
}

// 表名
func (SystemUser) TableName() string {
	return "system_user"
}
