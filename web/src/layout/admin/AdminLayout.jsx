import { useEffect, useState } from 'react';
import { App, Avatar, Badge, Dropdown, Layout, Menu } from 'antd';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { Logo } from '@/common/Image.jsx';
import { DynamicIcon } from '@/common/Icon.jsx';
import { GenerateGenderBadge } from '@/common/Tag.jsx';
import { FooterDescriptionComponent } from '@/common/Text.jsx';
import { RouteRules } from '@/route/RouteRules.jsx';
import { jwtDecode } from 'jwt-decode';
import { GetLocalToken } from '@/handler/Token.jsx';
import { AxiosGET } from '@/handler/Request.jsx';
import { BackendApiPrefix, BackendApiSuffix } from '@/common/Api.jsx';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { GenerateMenuTree } from '@/utils/Tree.jsx';

const { Header, Content, Sider } = Layout;

// 传递 pathname 从 RuleRules 中获取父级菜单列表
const findMenuParentPathList = (pathname) => {
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

  // 侧边菜单展开宽度
  const siderMenuWidth = 200;

  // 侧边菜单折叠宽度
  const siderMenuCollapsedWidth = 60;

  // 侧边菜单数据
  const [siderMenuTreeItems, setSiderMenuTreeItems] = useState([]);

  // 获取系统角色授权菜单列表的方法
  const getSystemRoleMenuListHandler = async () => {
    const systemRoleMenuListApi = BackendApiPrefix + BackendApiSuffix.System.Role.Auth.MenuList.Path;
    try {
      const res = await AxiosGET(systemRoleMenuListApi);
      if (res.code === 200) {
        SystemRoleStates.SystemRoleMenuList = res.data.list;
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('系统异常，请稍后再试');
    }
  };

  // 获取系统角色授权菜单列表
  useEffect(() => {
    const fetchAndGenerateMenu = async () => {
      await getSystemRoleMenuListHandler();
      const menuItems = GenerateMenuTree(0, SystemRoleStates.SystemRoleMenuList);
      setSiderMenuTreeItems(menuItems);
    };
    fetchAndGenerateMenu();
  }, []);

  // 监听 pathname 变化，更新菜单
  useEffect(() => {
    const keyList = findMenuParentPathList(pathname);
    setSiderMenuOpenKeys(keyList);
    setSiderMenuSelectedKeys(pathname);
  }, [pathname]);

  // 折叠展开侧边菜单处理函数
  const sideMenuCollapsedHandler = (value) => {
    if (value) {
      setSiderMenuOpenKeys([]);
    } else {
      setTimeout(() => {
        const keyList = findMenuParentPathList(pathname);
        setSiderMenuOpenKeys(keyList);
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
  // 当前用户信息
  const [currentUserInfo, setCurrentUserInfo] = useState({});

  // 用户注销，清理 Redis 缓存数据，清空本地 localStorage 数据，并跳转到登录页面
  const logoutHandler = async () => {
    const logoutApi = BackendApiPrefix + BackendApiSuffix.Public.Auth.Logout.Path;
    try {
      const res = await AxiosGET(logoutApi);
      if (res.code === 200) {
        message.success('注销成功');
        localStorage.clear();
        navigate('/login');
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('系统异常，请稍后再试');
    }
  };
  
  // 从 JWT Token 中解析出用户信息
  useEffect(() => {
    const token = GetLocalToken();
    if (token) {
      setCurrentUserInfo(jwtDecode(token));
    }
  }, []);

  // 顶部菜单栏下拉菜单，需要从 Token 中解析出用户相关信息
  const dropdownMenuItems = [
    { label: `${currentUserInfo?.cnName} / ${currentUserInfo?.enName}`, key: 'name', disabled: true },
    { label: <a target="_blank">修改资料</a>, key: 'me', onClick: () => navigate('/me') },
    { type: 'divider' },
    { label: '注销登录', key: 'logout', onClick: logoutHandler }
  ];

  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 获取用户授权接口列表
  /////////////////////////////////////////////////////////////////////////////////////////////////////////////
  // 获取系统角色授权接口列表的方法
  const getSystemRoleApiListHandler = async () => {
    const roleApiListApi = BackendApiPrefix + BackendApiSuffix.System.Role.Auth.ApiList.Path;
    try {
      const res = await AxiosGET(roleApiListApi);
      if (res.code === 200) {
        SystemRoleStates.SystemRoleApiList = res.data.list;
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('系统异常，请稍后再试');
    }
  };

  // 获取系统角色授权接口列表，主要用于按钮鉴权，全局使用
  useEffect(() => {
    getSystemRoleApiListHandler();
  }, []);

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
            <Badge count={GenerateGenderBadge(currentUserInfo?.gender)}>
              <Dropdown menu={{ items: dropdownMenuItems }}>
                <Avatar shape="circle" size={30} src={currentUserInfo?.avatar} />
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
              items={siderMenuTreeItems} // 菜单项
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
