import { ManOutlined, QuestionOutlined, WomanOutlined } from '@ant-design/icons';
import { Tag } from 'antd';

// 生成性别徽章
const GenerateGenderBadge = (gender) => {
  const icons = {
    1: <ManOutlined style={{ backgroundColor: '#165dff' }} />,
    2: <WomanOutlined style={{ backgroundColor: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ backgroundColor: '#999999' }} />
  };
  return icons[gender] || icons.default;
};

// 生成性别图标
const GenerateGenderIcon = (gender) => {
  const icons = {
    1: <ManOutlined style={{ color: '#165dff' }} />,
    2: <WomanOutlined style={{ color: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ color: '#999999' }} />
  };
  return icons[gender] || icons.default;
};

// 生成状态标签
const GenerateStatusTag = (status) => {
  const statusMap = {
    1: { text: '启用', color: 'green' },
    0: { text: '禁用', color: 'red' }
  };
  const { text, color } = statusMap[status] || {};
  return <Tag bordered={false} color={color}>{text}</Tag>;
};

// 生成角色标签
const GenerateRoleTag = (name) => {
  const roleColors = {
    超级管理员: 'magenta',
    管理员: 'volcano',
    读写: 'green',
    只读: 'warning'
  };
  const color = roleColors[name] || '';
  return <Tag bordered={false} color={color}>{name}</Tag>;
};

export { GenerateGenderBadge, GenerateGenderIcon, GenerateStatusTag, GenerateRoleTag };
