import { Navigate, useRoutes } from 'react-router';
import React from 'react';
import RouteLazyLoad from '@/route/RouteLazyLoad.jsx';

// 路由数据，参数规则：https://reactrouter.com/en/main/route/route
// 新增 auth 字段，用于校验该路由是开放路由还是需要认证才能查看的路由
export const RouteRules = [
  {
    path: '/', // 入口路由
    element: <Navigate to="/dashboard" />, // 路由跳转，默认跳转到其它页面
    auth: false // 用于认证
  },
  {
    path: '/', // 后台入口
    auth: true,
    element: RouteLazyLoad(React.lazy(() => import('../layout/admin/AdminLayout.jsx'))), // 模板继承
    children: [
      {
        path: '/dashboard',
        name: '工作空间',
        auth: true,
        element: RouteLazyLoad(React.lazy(() => import('../page/dashboard/Dashboard.jsx')))
      },
      {
        path: '/query',
        name: '即时查询',
        auth: true,
        element: RouteLazyLoad(React.lazy(() => import('../page/query/Query.jsx')))
      },
      {
        path: '/alert',
        name: '告警管理',
        auth: true,
        children: [
          {
            path: '/alert/active',
            name: '活跃告警',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/active/AlertActive.jsx')))
          },
          {
            path: '/alert/rule',
            name: '告警规则',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/rule/AlertRule.jsx')))
          },
          {
            path: '/alert/subscription',
            name: '告警订阅',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/subscription/AlertSubscription.jsx')))
          },
          {
            path: '/alert/block',
            name: '告警屏蔽',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/block/AlertBlock.jsx')))
          },
          {
            path: '/alert/history',
            name: '告警历史',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/history/AlertHistory.jsx')))
          },
          {
            path: '/alert/callback',
            name: '告警回调',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert/callback/AlertCallback.jsx')))
          }
        ]
      },
      {
        path: '/alert-notification',
        name: '告警通知',
        auth: true,
        children: [
          {
            path: '/alert-notification/media',
            name: '通知媒介',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert-notification/media/AlertNotificationMedia.jsx')))
          },
          {
            path: '/alert-notification/template',
            name: '通知模板',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert-notification/template/AlertNotificationTemplate.jsx')))
          }
        ]
      },
      {
        path: '/alert-group',
        name: '告警分组',
        auth: true,
        children: [
          {
            path: '/alert-group/project',
            name: '项目管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert-group/project/AlertGroupProject.jsx')))
          },
          {
            path: '/alert-group/team',
            name: '团队管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert-group/team/AlertGroupTeam.jsx')))
          },
          {
            path: '/alert-group/schedule',
            name: '人员排班',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/alert-group/schedule/AlertGroupSchedule.jsx')))
          }
        ]
      },
      {
        path: '/datasource',
        name: '数据来源',
        auth: true,
        element: RouteLazyLoad(React.lazy(() => import('../page/datasource/Datasource.jsx')))
      },
      {
        path: '/system',
        name: '系统管理',
        auth: true,
        children: [
          {
            path: '/system/department',
            name: '部门管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/department/SystemDepartment.jsx')))
          },
          {
            path: '/system/job-position',
            name: '职位管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/job-position/SystemJobPosition.jsx')))
          },
          {
            path: '/system/user',
            name: '用户管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/user/SystemUser.jsx')))
          },
          {
            path: '/system/role',
            name: '角色管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/role/SystemRole.jsx')))
          },
          {
            path: '/system/menu',
            name: '菜单管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/menu/SystemMenu.jsx')))
          },
          {
            path: '/system/api',
            name: '接口管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/api/SystemDepartment.jsx')))
          },
          {
            path: '/system/permission',
            name: '权限管理',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/system/permission/SystemPermission.jsx')))
          }
        ]
      },
      {
        path: '/me',
        name: '个人中心',
        auth: true,
        element: RouteLazyLoad(React.lazy(() => import('../page/me/Me.jsx')))
      },
      {
        path: '/security-audit',
        name: '安全审计',
        auth: true,
        children: [
          {
            path: '/security-audit/login',
            name: '登录日志',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/security-audit/login/SecurityAuditLogin.jsx')))
          },
          {
            path: '/security-audit/operation',
            name: '操作日志',
            auth: true,
            element: RouteLazyLoad(React.lazy(() => import('../page/security-audit/operation/SecurityAuditOperation.jsx')))
          }
        ]
      },
      {
        path: '/system-information',
        name: '系统信息',
        auth: true,
        element: RouteLazyLoad(React.lazy(() => import('../page/system-information/SystemInformation.jsx')))
      }
    ]
  },
  {
    path: '/', // 错误页面入口
    auth: false,
    element: RouteLazyLoad(React.lazy(() => import('../layout/error/ErrorLayout.jsx'))),
    children: [
      {
        path: '/error/403',
        name: '403',
        auth: false,
        element: RouteLazyLoad(React.lazy(() => import('../page/error/403.jsx')))
      },
      {
        path: '/error/404',
        name: '404',
        auth: false,
        element: RouteLazyLoad(React.lazy(() => import('../page/error/404.jsx')))
      },
      {
        path: '/error/500',
        name: '500',
        auth: false,
        element: RouteLazyLoad(React.lazy(() => import('../page/error/500.jsx')))
      }
    ]
  },
  {
    path: '/', // 登录页面入口
    auth: false,
    element: RouteLazyLoad(React.lazy(() => import('../layout/login/LoginLayout.jsx'))),
    children: [
      {
        path: '/login',
        name: '用户登录',
        auth: false,
        element: RouteLazyLoad(React.lazy(() => import('../page/login/Login.jsx')))
      },
      {
        path: '/reset',
        name: '密码重置',
        auth: false,
        element: RouteLazyLoad(React.lazy(() => import('../page/login/ResetPassword.jsx')))
      }
    ]
  },
  {
    path: '*', // 没有匹配默认路由
    element: <Navigate to="/error/404" />,
    auth: false
  }
];

// 生成可用的路由组件
export const Routes = () => {
  return useRoutes(RouteRules);
};
