package model

// 职位模型
type SystemJobPosition struct {
	BaseModel
	Name        string       `gorm:"uniqueIndex:idx_name;comment:职位名称" json:"name"`
	Creator     string       `gorm:"not null;comment:创建人，格式：username,cnName,enName,id" json:"creator"`
	SystemUsers []SystemUser `gorm:"many2many:system_user_job_position_relation;" json:"systemUsers,omitempty"`
}

// 表名
func (SystemJobPosition) TableName() string {
	return "system_job_position"
}
