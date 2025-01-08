package model

// 部门模型
type SystemDepartment struct {
	BaseModel
	ParentId    uint               `gorm:"uniqueIndex:uidx_parent_id_name;comment:父级部门ID" json:"parentId"`
	Name        string             `gorm:"uniqueIndex:uidx_parent_id_name;comment:部门名称" json:"name"` // 同一部门下子部门名称不能相同
	Creator     string             `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	Children    []SystemDepartment `gorm:"-" json:"children,omitempty"`
	SystemUsers []SystemUser       `gorm:"many2many:system_user_department_relation;" json:"systemUsers,omitempty"`
}

// 表名
func (SystemDepartment) TableName() string {
	return "system_department"
}
