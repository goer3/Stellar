// 存在 localStorage 中的参数定义
// Token
const TOKEN_KEY = 'token';
// Token 过期时间
const TOKEN_EXPIRE_TIME_KEY = 'tokenExpireTime';

// 校验 Token 是否过期
const TokenVerification = () => {
  const expireTime = localStorage.getItem(TOKEN_EXPIRE_TIME_KEY);
  if (!expireTime) {
    return false;
  }
  return new Date().getTime() < Date.parse(expireTime);
};

// 获取未过期的 Token
const GetUnexpireToken = () => {
  if (TokenVerification()) {
    const token = localStorage.getItem(TOKEN_KEY);
    // 验证 Token 是否为空
    if (token) {
      return token;
    }
  }
  // 清理 localStorage
  localStorage.clear();
  return null;
};

// 设置 Token 和过期时间
const SetTokenAndExpireTime = (token, expireTime) => {
  localStorage.setItem(TOKEN_KEY, token);
  localStorage.setItem(TOKEN_EXPIRE_TIME_KEY, expireTime);
};

// 清理 Token 和过期时间
const ClearTokenAndExpireTime = () => {
  localStorage.removeItem(TOKEN_KEY);
  localStorage.removeItem(TOKEN_EXPIRE_TIME_KEY);
};

export { TokenVerification, GetUnexpireToken, SetTokenAndExpireTime, ClearTokenAndExpireTime };
