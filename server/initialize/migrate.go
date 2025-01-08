package initialize

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 数据表同步
func MigrateTables() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行数据表同步")
	common.DB.AutoMigrate(&model.SystemRole{})
	common.DB.AutoMigrate(&model.SystemUser{})
	common.DB.AutoMigrate(&model.SystemMenu{})
	common.DB.AutoMigrate(&model.SystemJobPosition{})
	common.DB.AutoMigrate(&model.SystemDepartment{})
	common.DB.AutoMigrate(&model.SystemApiCategory{})
	common.DB.AutoMigrate(&model.SystemApi{})
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t数据表同步完成")
}
