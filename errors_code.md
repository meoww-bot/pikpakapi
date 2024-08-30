```json
{
  "error": "captcha_invalid",
  "error_code": 9,
  "error_url": "",
  "error_description": "验证码无效",
  "error_details": [
    {
      "@type": "type.googleapis.com/google.rpc.DebugInfo",
      "stack_entries": [],
      "detail": "failed to decode"
    }
  ]
}
```

```json
{
  "error": "unauthenticated",
  "error_code": 16,
  "error_url": "",
  "error_description": "帐号认证失败，请重新登录",
  "error_details": [
    {
      "@type": "type.googleapis.com/google.rpc.DebugInfo",
      "stack_entries": [],
      "detail": "oidc: token is expired (Token Expiry: 2024-08-29 20:10:14 +0800 CST)"
    }
  ]
}
```

```json
{
  "error": "captcha_invalid",
  "error_code": 4002,
  "error_description": "meta.username expect ****, but got map[captcha_sign:**** client_version:**** package_name:**** result:accept timestamp:**** user_id:], please check captcha init params",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.ErrorInfo",
      "reason": "meta.username expect ****, but got map[captcha_sign:**** client_version:**** package_name:**** result:accept timestamp:**** user_id:], please check captcha init params"
    },
    {
      "@type": "type.googleapis.com/google.rpc.LocalizedMessage",
      "locale": "zh",
      "message": "安全验证未通过，请重新验证"
    }
  ]
}
```

```json
{
  "error": "invalid_grant",
  "error_code": 4126,
  "error_description": "refresh token **** has been refresh at 2024-08-29T20:41:24+08:00",
  "details": [
    {
      "@type": "type.googleapis.com/google.rpc.ErrorInfo",
      "reason": "refresh token **** has been refresh at 2024-08-29T20:41:24+08:00"
    },
    {
      "@type": "type.googleapis.com/google.rpc.LocalizedMessage",
      "locale": "zh",
      "message": "用户凭证不正确或已过期"
    }
  ]
}
```
