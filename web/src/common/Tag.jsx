import { ManOutlined, QuestionOutlined, WomanOutlined } from '@ant-design/icons';

// 生成性别徽章
export const generateGenderBadge = (gender) => {
  return gender === 1 ? (
    <ManOutlined style={{ backgroundColor: '#165dff' }} />
  ) : gender === 2 ? (
    <WomanOutlined style={{ backgroundColor: '#ff4d4f' }} />
  ) : (
    <QuestionOutlined style={{ backgroundColor: '#999999' }} />
  );
};
