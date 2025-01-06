import {
  AlertOutlined,
  ApiOutlined,
  ClusterOutlined,
  ConsoleSqlOutlined,
  DeploymentUnitOutlined,
  DesktopOutlined,
  FileProtectOutlined,
  InsuranceOutlined,
  MailOutlined,
  ProjectOutlined,
  SettingOutlined,
  UsergroupAddOutlined,
  UserOutlined
} from '@ant-design/icons';

// 图标字符串映射
export const IconMap = {
  DesktopOutlined: DesktopOutlined,
  ConsoleSqlOutlined: ConsoleSqlOutlined,
  ApiOutlined: ApiOutlined,
  SettingOutlined: SettingOutlined,
  UsergroupAddOutlined: UsergroupAddOutlined,
  InsuranceOutlined: InsuranceOutlined,
  UserOutlined: UserOutlined,
  FileProtectOutlined: FileProtectOutlined,
  ClusterOutlined: ClusterOutlined,
  AlertOutlined: AlertOutlined,
  MailOutlined: MailOutlined,
  ProjectOutlined: ProjectOutlined,
  DeploymentUnitOutlined: DeploymentUnitOutlined
};

// 生成 Icon
// 用法：<DynamicIcon iconName={'DesktopOutlined'} />
export const DynamicIcon = ({ iconName }) => {
  const IconComponent = IconMap[iconName];
  return IconComponent ? <IconComponent /> : null;
};
