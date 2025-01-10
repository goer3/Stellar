package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 系统API分类数据
var systemApiCategories = []model.SystemApiCategory{}

// 迁移系统API分类数据
func MigrateSystemApiCategoryData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行系统API分类数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_api_category")
	common.DB.Create(&systemApiCategories)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t系统API分类数据迁移完成")
}
