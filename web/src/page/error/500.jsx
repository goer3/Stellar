import { BG500 } from '@/common/Image.jsx';

const PageInternalServerError = () => {
  return (
    <>
      <img className="admin-unselect" src={BG500} alt="" />
      <div className="admin-error-code admin-unselect">500</div>
    </>
  );
};

export default PageInternalServerError;
