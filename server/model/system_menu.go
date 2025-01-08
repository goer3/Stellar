package model

// 菜单模型
type SystemMenu struct {
	BaseModel
	ParentId    uint         `gorm:"uniqueIndex:uidx_parent_id_label_key;comment:父级菜单ID" json:"parentId"` // 三个字段联合唯一
	Label       string       `gorm:"uniqueIndex:uidx_parent_id_label_key;comment:菜单名称" json:"label"`
	Key         string       `gorm:"uniqueIndex:uidx_parent_id_label_key;comment:菜单路径" json:"key"`
	Icon        string       `gorm:"comment:菜单图标" json:"icon"`
	Sort        uint         `gorm:"comment:排序" json:"sort"`
	Creator     string       `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	Children    []SystemMenu `gorm:"-" json:"children,omitempty"`
	SystemRoles []SystemRole `gorm:"many2many:system_role_menu_relation;" json:"systemRoles,omitempty"`
}

// 表名
func (SystemMenu) TableName() string {
	return "system_menu"
}
