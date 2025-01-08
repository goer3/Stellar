package model

// 接口分类模型
type SystemApiCategory struct {
	BaseModel
	Name       string      `gorm:"uniqueIndex:idx_name;comment:接口分类名称" json:"name"`
	SystemApis []SystemApi `gorm:"-" json:"systemApis,omitempty"`
}

// 表名
func (SystemApiCategory) TableName() string {
	return "system_api_category"
}
