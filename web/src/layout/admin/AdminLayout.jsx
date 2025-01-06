import { useEffect, useState } from 'react';
import { App, Avatar, Badge, Dropdown, Layout, Menu } from 'antd';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { jwtDecode } from 'jwt-decode';
import { Logo } from '@/common/Image.jsx';
import { FooterText } from '@/common/Text.jsx';
import { TreeFindPath } from '@/utils/Path.jsx';
import { RouteRules } from '@/routes/RouteRules.jsx';
import { DynamicIcon } from '@/utils/IconLoad.jsx';
import { AxiosGet } from '@/handler/HandlerUtilsRequest.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import { ManOutlined, QuestionOutlined, WomanOutlined } from '@ant-design/icons';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';

const { Header, Content, Sider } = Layout;

// 生成菜单
const getItem = (label, key, icon, children) => ({
  key,
  icon,
  children,
  label
});

// 侧边菜单
const siderMenus = [
  getItem('工作空间', '/dashboard', <DynamicIcon iconName={'DesktopOutlined'} />),
  getItem('即时查询', '/query', <DynamicIcon iconName={'ConsoleSqlOutlined'} />),
  getItem('告警管理', '/alert', <DynamicIcon iconName={'AlertOutlined'} />, [
    getItem('活跃告警', '/alert/active'),
    getItem('告警规则', '/alert/rule'),
    getItem('告警订阅', '/alert/subscription'),
    getItem('告警屏蔽', '/alert/block'),
    getItem('告警历史', '/alert/history'),
    getItem('告警回调', '/alert/callback')
  ]),
  getItem('告警通知', '/alert-notification', <DynamicIcon iconName={'MailOutlined'} />, [getItem('通知媒介', '/alert-notification/media'), getItem('通知模板', '/alert-notification/template')]),
  getItem('告警分组', '/alert-group', <DynamicIcon iconName={'ProjectOutlined'} />, [
    getItem('项目管理', '/alert-group/project'),
    getItem('团队管理', '/alert-group/team'),
    getItem('人员排班', '/alert-group/schedule')
  ]),
  getItem('数据来源', '/datasource', <DynamicIcon iconName={'ApiOutlined'} />),
  getItem('系统设置', '/system', <DynamicIcon iconName={'ClusterOutlined'} />, [
    getItem('部门管理', '/system/department'),
    getItem('职位管理', '/system/job-position'),
    getItem('用户管理', '/system/user'),
    getItem('角色管理', '/system/role'),
    getItem('菜单管理', '/system/menu'),
    getItem('接口管理', '/system/api'),
    getItem('权限管理', '/system/permission')
  ]),
  getItem('个人中心', '/me', <DynamicIcon iconName={'UserOutlined'} />),
  getItem('安全审计', '/security-audit', <DynamicIcon iconName={'FileProtectOutlined'} />, [getItem('登录日志', '/security-audit/login'), getItem('操作日志', '/security-audit/operation')]),
  getItem('系统信息', '/system-information', <DynamicIcon iconName={'DeploymentUnitOutlined'} />)
];

const AdminLayout = () => {
  // 消息提示
  const { message } = App.useApp();
  // 菜单跳转
  const navigate = useNavigate();
  // 菜单展开收起状态
  const [collapsed, setCollapsed] = useState(false);
  // 展开和收缩菜单宽度
  const menuWidth = 200;
  const menuCollapsedWidth = 60;

  // 获取当前的请求路径，并监听该路径是否改变，如果改变则修改页面菜单数据
  const { pathname } = useLocation(); // 当前页面
  const [openKeys, setOpenKeys] = useState([pathname]); // 展开菜单，父级菜单
  const [selectedKeys, setSelectedKeys] = useState([pathname]); // 选中菜单
  const [rememberedOpenKeys, setRememberedOpenKeys] = useState([]); // 添加一个状态来保存展开的菜单keys
  useEffect(() => {
    const initialPath = TreeFindPath(RouteRules, (data) => data.path === pathname);
    setOpenKeys(initialPath);
    setSelectedKeys(pathname);
  }, []);

  // 修改折叠处理函数，添加延迟恢复展开菜单的功能
  const handleCollapse = (value) => {
    if (value) {
      setRememberedOpenKeys(openKeys);
      setOpenKeys([]);
    } else {
      // 添加 100ms 延迟，避免动画卡顿
      setTimeout(() => {
        setOpenKeys(rememberedOpenKeys);
      }, 100);
    }
    setCollapsed(value);
  };

  // 添加菜单点击处理函数
  const handleMenuClick = ({ key }) => {
    setSelectedKeys([key]);
    navigate(key);
  };

  // 用户注销
  const logoutHandler = async () => {
    const res = await AxiosGet(Apis.Public.Logout);
    if (res.code === 200) {
      localStorage.clear();
      message.success('注销成功');
      navigate('/login');
    } else {
      message.error('注销异常，' + res.message);
    }
  };

  // 从 Token 中解析出用户相关信息
  const token = localStorage.getItem('token');
  const userInfo = jwtDecode(token);

  // 下拉菜单
  const dropdownMenus = [
    {
      label: `${userInfo?.cnName} / ${userInfo?.enName}`,
      key: '0',
      disabled: true
    },
    {
      label: <a target="_blank">修改资料</a>,
      key: '1',
      onClick: () => navigate('/me')
    },
    {
      type: 'divider'
    },
    {
      label: '注销登录',
      key: '2',
      onClick: logoutHandler
    }
  ];

  // 获取角色API列表
  useEffect(() => {
    const getSystemRoleApis = async () => {
      try {
        const res = await AxiosGet(Apis.System.Role.ApiList);
        if (res.code === 200) {
          SystemRoleStates.SystemRoleApis = res.data;
        } else {
          message.error(res.message);
        }
      } catch (error) {
        message.error('获取角色API列表异常');
      }
    };
    getSystemRoleApis();
  }, []);

  return (
    <Layout>
      <Header className="admin-header">
        <div className="admin-left">
          <div className="admin-logo">
            <img className="admin-unselect" src={Logo} alt="" />
          </div>
        </div>
        <div className="admin-right">
          <Badge
            count={
              userInfo?.gender === 1 ? (
                <ManOutlined style={{ backgroundColor: '#165dff' }} />
              ) : userInfo?.gender === 2 ? (
                <WomanOutlined style={{ backgroundColor: '#ff4d4f' }} />
              ) : (
                <QuestionOutlined style={{ backgroundColor: '#999' }} />
              )
            }
          >
            <Dropdown
              menu={{
                items: dropdownMenus
              }}
            >
              <Avatar shape="circle" size={30} src={userInfo?.avatar} />
            </Dropdown>
          </Badge>
        </div>
      </Header>
      <Layout className="admin-main">
        <Sider className="admin-sider" theme="light" width={menuWidth} collapsedWidth={menuCollapsedWidth} collapsible collapsed={collapsed} onCollapse={handleCollapse}>
          <Menu className="admin-menu" mode="inline" openKeys={openKeys} onOpenChange={setOpenKeys} selectedKeys={selectedKeys} items={siderMenus} onClick={handleMenuClick} />
        </Sider>
        <Layout className="admin-body">
          <Content className="admin-content">
            <Outlet />
          </Content>
          <div className="admin-footer">
            <FooterText />
          </div>
        </Layout>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;
