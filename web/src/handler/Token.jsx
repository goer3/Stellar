// 存在 localStorage 中的参数定义
// Token
const TOKEN_KEY = 'token';
// Token 过期时间
const TOKEN_EXPIRE_TIME_KEY = 'tokenExpireTime';

// 本地校验 Token 是否过期
const VerifyLocalTokenExpireTime = () => {
  const expireTime = localStorage.getItem(TOKEN_EXPIRE_TIME_KEY);
  if (!expireTime) {
    return false;
  }
  return new Date().getTime() < Date.parse(expireTime);
};

// 获取本地 Token
const GetLocalToken = () => {
  const token = localStorage.getItem(TOKEN_KEY);
  if (token) {
    return token;
  }
  return null;
};

// 设置 Token 和过期时间
const SetLocalTokenAndExpireTime = (token, expireTime) => {
  localStorage.setItem(TOKEN_KEY, token);
  localStorage.setItem(TOKEN_EXPIRE_TIME_KEY, expireTime);
};

// 清理 Token 和过期时间
const ClearLocalTokenAndExpireTime = () => {
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(TOKEN_EXPIRE_TIME_KEY);
};

export { VerifyLocalTokenExpireTime, GetLocalToken, SetLocalTokenAndExpireTime, ClearLocalTokenAndExpireTime };
