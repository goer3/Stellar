import { Outlet, useNavigate } from 'react-router';
import { Button } from 'antd';
import { ArrowLeftOutlined } from '@ant-design/icons';
import { Logo } from '@/common/Image.jsx';
import { FooterDescriptionComponent } from '@/common/Text.jsx';

const ErrorLayout = () => {
  const navigate = useNavigate();
  return (
    <>
      <div className="admin-error">
        <div className="admin-error-header">
          <img className="admin-unselect" src={Logo} alt="logo" />
        </div>
        <div className="admin-error-body">
          <Outlet />
          <Button type="primary" icon={<ArrowLeftOutlined />} onClick={() => navigate('/')}>
            回到首页
          </Button>
        </div>
        <div className="admin-error-footer admin-unselect">
          <FooterDescriptionComponent />
        </div>
      </div>
    </>
  );
};

export default ErrorLayout;
