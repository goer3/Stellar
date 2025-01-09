package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 部门数据
var systemDepartments = []model.SystemDepartment{
	{
		BaseModel: model.BaseModel{Id: 1},
		ParentId:  0,
		Name:      "某某科技公司",
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemUsers: []model.SystemUser{
			systemUsers[0],
		},
		Children: []model.SystemDepartment{
			{
				BaseModel: model.BaseModel{Id: 2},
				ParentId:  1,
				Name:      "产品中心",
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemUsers: []model.SystemUser{
					systemUsers[0],
				},
				Children: []model.SystemDepartment{
					{
						BaseModel: model.BaseModel{Id: 3},
						ParentId:  2,
						Name:      "产品部",
						Creator:   common.SYSTEM_DEFAULT_CREATOR,
					},
					{
						BaseModel: model.BaseModel{Id: 4},
						ParentId:  2,
						Name:      "设计部",
						Creator:   common.SYSTEM_DEFAULT_CREATOR,
					},
				},
			},
			{
				BaseModel: model.BaseModel{Id: 5},
				ParentId:  1,
				Name:      "研发中心",
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemUsers: []model.SystemUser{
					systemUsers[1],
				},
				Children: []model.SystemDepartment{
					{
						BaseModel: model.BaseModel{Id: 6},
						ParentId:  5,
						Name:      "研发部",
						Creator:   common.SYSTEM_DEFAULT_CREATOR,
						Children: []model.SystemDepartment{
							{
								BaseModel: model.BaseModel{Id: 7},
								ParentId:  6,
								Name:      "前端开发部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 8},
								ParentId:  6,
								Name:      "后端开发部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
						},
					},
					{
						BaseModel: model.BaseModel{Id: 9},
						ParentId:  5,
						Name:      "测试部",
						Creator:   common.SYSTEM_DEFAULT_CREATOR,
						SystemUsers: []model.SystemUser{
							systemUsers[3],
						},
					},
					{
						BaseModel: model.BaseModel{Id: 10},
						ParentId:  5,
						Name:      "运维部",
						Creator:   common.SYSTEM_DEFAULT_CREATOR,
						Children: []model.SystemDepartment{
							{
								BaseModel: model.BaseModel{Id: 11},
								ParentId:  10,
								Name:      "系统运维部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 12},
								ParentId:  10,
								Name:      "网络运维部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 13},
								ParentId:  10,
								Name:      "安全运维部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 14},
								ParentId:  10,
								Name:      "数据库运维部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 15},
								ParentId:  10,
								Name:      "业务运维部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
							},
							{
								BaseModel: model.BaseModel{Id: 16},
								ParentId:  10,
								Name:      "运维开发部",
								Creator:   common.SYSTEM_DEFAULT_CREATOR,
								SystemUsers: []model.SystemUser{
									systemUsers[2],
								},
							},
						},
					},
				},
			},
		},
	},
}

// 递归插入数据方法
func insertSystemDepartmentsData(departments []model.SystemDepartment) {
	for _, department := range departments {
		common.DB.Create(&department)
		if len(department.Children) > 0 {
			insertSystemDepartmentsData(department.Children)
		}
	}
}

// 迁移部门数据
func MigrateSystemDepartmentData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行部门数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_department")
	insertSystemDepartmentsData(systemDepartments)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t部门数据迁移完成")
}
