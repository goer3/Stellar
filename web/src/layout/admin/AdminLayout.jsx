import { useEffect, useState } from 'react';
import { App, Avatar, Badge, Dropdown, Layout, Menu } from 'antd';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { Logo } from '@/common/Image.jsx';
import { DynamicIcon } from '@/common/Icon.jsx';
import { generateGenderBadge } from '@/common/Tag.jsx';
import { FooterDescriptionComponent } from '@/common/Text.jsx';
import { RouteRules } from '@/route/RouteRules.jsx';

const { Header, Content, Sider } = Layout;

// 生成菜单的方法
const generateMenuItem = (label, key, icon, children) => ({
  key,
  icon: <DynamicIcon iconName={icon} />,
  children,
  label
});

// 传递 pathname 从 RuleRules 中获取父级菜单列表
export const findMenuParentPathList = (pathname) => {
  if (!pathname) return [];
  const findPath = (items, targetPath, parentPath = []) => {
    for (const item of items) {
      const currentPath = [...parentPath, item.path];
      if (item.path === targetPath) {
        return currentPath;
      }
      if (item.children?.length) {
        const foundPath = findPath(item.children, targetPath, currentPath);
        if (foundPath.length) {
          return foundPath;
        }
      }
    }
    return [];
  };
  return findPath(RouteRules, pathname);
};

/////////////////////////////////////////////////////////////////////////////////////////////////////////////
// AdminLayout 组件
/////////////////////////////////////////////////////////////////////////////////////////////////////////////
const AdminLayout = () => {
  // 消息提示
  const { message } = App.useApp();

  // 菜单跳转
  const navigate = useNavigate();

  // 当前路径，并需要监听该路径是否改变，如果改变则需要渲染对应页面
  const { pathname } = useLocation();

  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 侧边菜单
  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 侧边菜单是否折叠
  const [siderMenuCollapsed, setSiderMenuCollapsed] = useState(false);

  // 展开菜单
  const [siderMenuOpenKeys, setSiderMenuOpenKeys] = useState([pathname]);

  // 选中菜单
  const [siderMenuSelectedKeys, setSiderMenuSelectedKeys] = useState([pathname]);

  // 保存展开的菜单 keys，用于确保展开收起状态切换还能继续展开
  const [siderMenuRememberedOpenKeys, setSiderMenuRememberedOpenKeys] = useState([]);

  // 侧边菜单展开宽度
  const siderMenuWidth = 200;

  // 侧边菜单折叠宽度
  const siderMenuCollapsedWidth = 60;

  // 侧边菜单数据
  const siderMenuItems = [
    generateMenuItem('工作空间', '/dashboard', 'DesktopOutlined'),
    generateMenuItem('即时查询', '/query', 'ConsoleSqlOutlined'),
    generateMenuItem('告警管理', '/alert', 'AlertOutlined', [
      generateMenuItem('活跃告警', '/alert/active'),
      generateMenuItem('告警规则', '/alert/rule'),
      generateMenuItem('告警订阅', '/alert/subscription'),
      generateMenuItem('告警屏蔽', '/alert/block'),
      generateMenuItem('告警历史', '/alert/history'),
      generateMenuItem('告警回调', '/alert/callback')
    ]),
    generateMenuItem('告警通知', '/alert-notification', 'MailOutlined', [generateMenuItem('通知媒介', '/alert-notification/media'), generateMenuItem('通知模板', '/alert-notification/template')]),
    generateMenuItem('告警分组', '/alert-group', 'ProjectOutlined', [
      generateMenuItem('项目管理', '/alert-group/project'),
      generateMenuItem('团队管理', '/alert-group/team'),
      generateMenuItem('人员排班', '/alert-group/schedule')
    ]),
    generateMenuItem('数据来源', '/datasource', 'ApiOutlined'),
    generateMenuItem('系统设置', '/system', 'ClusterOutlined', [
      generateMenuItem('部门管理', '/system/department'),
      generateMenuItem('职位管理', '/system/job-position'),
      generateMenuItem('用户管理', '/system/user'),
      generateMenuItem('角色管理', '/system/role'),
      generateMenuItem('菜单管理', '/system/menu'),
      generateMenuItem('接口管理', '/system/api'),
      generateMenuItem('权限管理', '/system/permission')
    ]),
    generateMenuItem('个人中心', '/me', 'UserOutlined'),
    generateMenuItem('安全审计', '/security-audit', 'FileProtectOutlined', [generateMenuItem('登录日志', '/security-audit/login'), generateMenuItem('操作日志', '/security-audit/operation')]),
    generateMenuItem('系统信息', '/system-information', 'DeploymentUnitOutlined')
  ];

  // 监听 pathname 变化，更新菜单
  useEffect(() => {
    const keyList = findMenuParentPathList(pathname);
    setSiderMenuOpenKeys(keyList);
    setSiderMenuSelectedKeys(pathname);
  }, [pathname]);

  // 折叠展开侧边菜单处理函数
  const sideMenuCollapsedHandler = (value) => {
    if (value) {
      // 记住选中菜单，避免状态切换后不继续选中和展开菜单
      setSiderMenuRememberedOpenKeys(siderMenuOpenKeys);
      setSiderMenuOpenKeys([]);
    } else {
      // 添加 100ms 延迟，避免动画卡顿
      setTimeout(() => {
        setSiderMenuOpenKeys(siderMenuRememberedOpenKeys);
      }, 100);
    }
    setSiderMenuCollapsed(value);
  };

  // 点击侧边菜单处理函数
  const siderMenuClickHandler = ({ key }) => {
    setSiderMenuSelectedKeys([key]);
    navigate(key);
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 顶部菜单栏
  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 用户注销，清理 Redis 缓存数据，清空本地 localStorage 数据，并跳转到登录页面
  const logoutHandler = async () => {
    console.log('注销');
  };

  // 顶部菜单栏下拉菜单，需要从 Token 中解析出用户相关信息
  const dropdownMenuItems = [
    { label: `吴彦祖 / DK`, key: '0', disabled: true },
    { label: <a target="_blank">修改资料</a>, key: '1', onClick: () => navigate('/me') },
    { type: 'divider' },
    { label: '注销登录', key: '2', onClick: logoutHandler }
  ];

  return (
    <>
      <Layout>
        <Header className="admin-header">
          <div className="admin-left">
            <div className="admin-logo">
              <img className="admin-unselect" src={Logo} alt="" />
            </div>
          </div>
          <div className="admin-right">
            <Badge count={generateGenderBadge(1)}>
              <Dropdown menu={{ items: dropdownMenuItems }}>
                <Avatar shape="circle" size={30} src="/image/avatar/default.png" />
              </Dropdown>
            </Badge>
          </div>
        </Header>
        <Layout className="admin-main">
          <Sider
            className="admin-sider"
            theme="light" // 主题
            width={siderMenuWidth} // 不折叠宽度
            collapsible // 是否可折叠
            collapsedWidth={siderMenuCollapsedWidth} // 折叠宽度
            collapsed={siderMenuCollapsed} // 是否折叠状态
            onCollapse={sideMenuCollapsedHandler} // 折叠展开处理函数
          >
            <Menu
              className="admin-menu"
              mode="inline" // 菜单模式
              openKeys={siderMenuOpenKeys} // 展开菜单
              onOpenChange={setSiderMenuOpenKeys} // 展开菜单改变处理函数
              selectedKeys={siderMenuSelectedKeys} // 选中菜单
              items={siderMenuItems} // 菜单项
              onClick={siderMenuClickHandler} // 点击菜单处理函数
            />
          </Sider>
          <Layout className="admin-body">
            <Content className="admin-content">
              <Outlet />
            </Content>
            <div className="admin-footer">
              <FooterDescriptionComponent />
            </div>
          </Layout>
        </Layout>
      </Layout>
    </>
  );
};

export default AdminLayout;
