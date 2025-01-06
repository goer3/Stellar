import { BG404 } from '@/common/Image.jsx';

const PageNotFoundError = () => {
  return (
    <>
      <img className="admin-unselect" src={BG404} alt="" />
      <div className="admin-error-code admin-unselect">404</div>
    </>
  );
};

export default PageNotFoundError;
