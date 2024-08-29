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
