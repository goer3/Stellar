package model

// 接口模型
type SystemApi struct {
	BaseModel
	Name                string             `gorm:"uniqueIndex:uidx_name_method_path;comment:接口名称" json:"name"` // 联合唯一
	Method              string             `gorm:"uniqueIndex:uidx_name_method_path;comment:请求方法" json:"method"`
	Path                string             `gorm:"uniqueIndex:uidx_name_method_path;comment:请求路径" json:"path"`
	Description         string             `gorm:"comment:接口描述" json:"description"`
	Permission          *uint              `gorm:"default:1;comment:权限标识，0：无需权限，1：需要权限" json:"permission"`
	Disabled            *uint              `gorm:"default:0;comment:是否禁止授权，0：否，1：是" json:"disabled"`
	Creator             string             `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	SystemApiCategoryId uint               `gorm:"comment:接口分类ID" json:"systemApiCategoryId"`
	SystemApiCategory   *SystemApiCategory `gorm:"foreignKey:SystemApiCategoryId;" json:"systemApiCategory,omitempty"`
	SystemRoles         []SystemRole       `gorm:"many2many:system_role_api_relation;" json:"systemRoles,omitempty"`
}

// 表名
func (SystemApi) TableName() string {
	return "system_api"
}
