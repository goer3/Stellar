import { Outlet } from 'react-router';
import { Logo } from '@/common/Image.jsx';
import { FooterDescriptionComponent } from '@/common/Text.jsx';

const LoginLayout = () => {
  return (
    <>
      <div className="admin-login-container">
        <div className="admin-login-box">
          <div className="admin-login-header">
            <img className="admin-unselect" src={Logo} alt="logo" />
          </div>
          <Outlet />
          <div className="admin-login-footer admin-unselect">
            <FooterDescriptionComponent />
          </div>
        </div>
      </div>
    </>
  );
};

export default LoginLayout;
