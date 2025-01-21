package model

// 系统设置
type SystemSetting struct {
	BaseModel
	Key         string `gorm:"not null;uniqueIndex:idx_key;comment:键" json:"key"`
	Value       string `gorm:"not null;comment:值" json:"value"`
	Description string `gorm:"comment:描述" json:"description"`
	Updater     string `gorm:"not null;comment:更新人" json:"updater"`
}

func (SystemSetting) TableName() string {
	return "system_setting"
}
