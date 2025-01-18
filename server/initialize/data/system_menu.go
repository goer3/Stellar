package data

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"time"
)

// 菜单数据
var systemMenus = []model.SystemMenu{
	{
		BaseModel:   model.BaseModel{Id: 1},
		ParentId:    0,
		Label:       "工作空间",
		Key:         "/dashboard",
		Icon:        "DesktopOutlined",
		Sort:        1,
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: systemRoles,
	},
	{
		BaseModel:   model.BaseModel{Id: 1000},
		ParentId:    0,
		Label:       "即时查询",
		Key:         "/query",
		Icon:        "ConsoleSqlOutlined",
		Sort:        1000,
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: systemRoles,
	},
	{
		BaseModel: model.BaseModel{Id: 2000},
		ParentId:  0,
		Label:     "告警管理",
		Key:       "/alert",
		Icon:      "AlertOutlined",
		Sort:      2000,
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				BaseModel: model.BaseModel{Id: 2100},
				ParentId:  2000,
				Label:     "活跃告警",
				Key:       "/alert/active",
				Icon:      "",
				Sort:      2100,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 2200},
				ParentId:  2000,
				Label:     "告警规则",
				Key:       "/alert/rule",
				Icon:      "",
				Sort:      2200,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 2300},
				ParentId:  2000,
				Label:     "告警订阅",
				Key:       "/alert/subscription",
				Icon:      "",
				Sort:      2300,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 2400},
				ParentId:  2000,
				Label:     "告警屏蔽",
				Key:       "/alert/block",
				Icon:      "",
				Sort:      2400,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 2500},
				ParentId:  2000,
				Label:     "告警历史",
				Key:       "/alert/history",
				Icon:      "",
				Sort:      2500,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 2600},
				ParentId:  2000,
				Label:     "告警回调",
				Key:       "/alert/callback",
				Icon:      "",
				Sort:      2600,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 3000},
		ParentId:  0,
		Label:     "告警通知",
		Key:       "/alert-notification",
		Icon:      "MailOutlined",
		Sort:      3000,
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				BaseModel: model.BaseModel{Id: 3100},
				ParentId:  3000,
				Label:     "通知媒介",
				Key:       "/alert-notification/media",
				Icon:      "",
				Sort:      3100,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 3200},
				ParentId:  3000,
				Label:     "通知模板",
				Key:       "/alert-notification/template",
				Icon:      "",
				Sort:      3200,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 4000},
		ParentId:  0,
		Label:     "告警分组",
		Key:       "/alert-group",
		Icon:      "ProjectOutlined",
		Sort:      4000,
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				BaseModel: model.BaseModel{Id: 4100},
				ParentId:  4000,
				Label:     "项目管理",
				Key:       "/alert-group/project",
				Icon:      "",
				Sort:      4100,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 4200},
				ParentId:  4000,
				Label:     "团队管理",
				Key:       "/alert-group/team",
				Icon:      "",
				Sort:      4200,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 4300},
				ParentId:  4000,
				Label:     "人员排班",
				Key:       "/alert-group/schedule",
				Icon:      "",
				Sort:      4300,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
			},
		},
	},
	{
		BaseModel: model.BaseModel{Id: 5000},
		ParentId:  0,
		Label:     "数据来源",
		Key:       "/datasource",
		Icon:      "ApiOutlined",
		Sort:      5000,
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
	},
	{
		BaseModel: model.BaseModel{Id: 6000},
		ParentId:  0,
		Label:     "系统设置",
		Key:       "/system",
		Icon:      "ClusterOutlined",
		Sort:      6000,
		Creator:   common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				BaseModel: model.BaseModel{Id: 6100},
				ParentId:  6000,
				Label:     "部门管理",
				Key:       "/system/department",
				Icon:      "",
				Sort:      6100,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6200},
				ParentId:  6000,
				Label:     "职位管理",
				Key:       "/system/job-position",
				Icon:      "",
				Sort:      6200,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6300},
				ParentId:  6000,
				Label:     "用户管理",
				Key:       "/system/user",
				Icon:      "",
				Sort:      6300,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6400},
				ParentId:  6000,
				Label:     "角色管理",
				Key:       "/system/role",
				Icon:      "",
				Sort:      6400,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6500},
				ParentId:  6000,
				Label:     "菜单管理",
				Key:       "/system/menu",
				Icon:      "",
				Sort:      6500,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6600},
				ParentId:  6000,
				Label:     "接口管理",
				Key:       "/system/api",
				Icon:      "",
				Sort:      6600,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				BaseModel: model.BaseModel{Id: 6700},
				ParentId:  6000,
				Label:     "权限管理",
				Key:       "/system/permission",
				Icon:      "",
				Sort:      6700,
				Creator:   common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		BaseModel:   model.BaseModel{Id: 7000},
		ParentId:    0,
		Label:       "个人中心",
		Key:         "/me",
		Icon:        "UserOutlined",
		Sort:        7000,
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: systemRoles,
	},
	{
		BaseModel:   model.BaseModel{Id: 8000},
		ParentId:    0,
		Label:       "安全审计",
		Key:         "/security-audit",
		Icon:        "FileProtectOutlined",
		Sort:        8000,
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: systemRoles,
		Children: []model.SystemMenu{
			{
				BaseModel:   model.BaseModel{Id: 8100},
				ParentId:    8000,
				Label:       "登录日志",
				Key:         "/security-audit/login",
				Icon:        "",
				Sort:        8100,
				Creator:     common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: systemRoles,
			},
			{
				BaseModel:   model.BaseModel{Id: 8200},
				ParentId:    8000,
				Label:       "操作日志",
				Key:         "/security-audit/operation",
				Icon:        "",
				Sort:        8200,
				Creator:     common.SYSTEM_DEFAULT_CREATOR,
				SystemRoles: systemRoles,
			},
		},
	},
	{
		BaseModel:   model.BaseModel{Id: 9000},
		ParentId:    0,
		Label:       "系统信息",
		Key:         "/system-information",
		Icon:        "DeploymentUnitOutlined",
		Sort:        9000,
		Creator:     common.SYSTEM_DEFAULT_CREATOR,
		SystemRoles: systemRoles,
	},
}

// 递归插入数据方法
func insertSystemMenusData(menus []model.SystemMenu) {
	for _, menu := range menus {
		common.DB.Create(&menu)
		if len(menu.Children) > 0 {
			insertSystemMenusData(menu.Children)
		}
	}
}

// 迁移菜单数据
func MigrateSystemMenuData() {
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t开始进行菜单数据迁移")
	common.DB.Exec("TRUNCATE TABLE system_menu")
	insertSystemMenusData(systemMenus)
	fmt.Println(time.Now().Format(common.TIME_FORMAT_MILLISECOND), "\t菜单数据迁移完成")
}
