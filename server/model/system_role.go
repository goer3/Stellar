package model

// 角色模型
type SystemRole struct {
	BaseModel
	Name        string       `gorm:"uniqueIndex:idx_name;comment:角色名称" json:"name"`
	Keyword     string       `gorm:"uniqueIndex:idx_keyword;comment:角色关键字" json:"keyword"`
	Description string       `gorm:"not null;comment:角色描述" json:"description"`
	Creator     string       `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	SystemUsers []SystemUser `gorm:"-" json:"systemUsers,omitempty"`
	SystemMenus []SystemMenu `gorm:"many2many:system_role_menu_relation;" json:"systemMenus,omitempty"`
	SystemApis  []SystemApi  `gorm:"many2many:system_role_api_relation;" json:"systemApis,omitempty"`
}

// 表名
func (SystemRole) TableName() string {
	return "system_role"
}
