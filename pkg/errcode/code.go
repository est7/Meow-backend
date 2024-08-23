package errcode

// 预定义错误
// 1 为系统级错误；2 为普通错误，

var (
	Success                     = NewCustomError(0, "Ok")
	ErrInternalServer           = NewCustomError(10001, "Internal server error")                         // 服务器内部错误
	ErrServiceUnavailable       = NewCustomError(10002, "Service unavailable")                           // 服务暂停
	ErrInvalidParam             = NewCustomError(10003, "Invalid params")                                // 参数错误
	ErrUnauthorized             = NewCustomError(10004, "Unauthorized error")                            // 未授权
	ErrNotFound                 = NewCustomError(10005, "Resource Not found")                            // 资源未找到
	ErrApiNotFound              = NewCustomError(10006, "Api not found")                                 // 接口不存在
	ErrDeadlineExceeded         = NewCustomError(10007, "RPC Deadline exceeded")                         // RPC 调用超时
	ErrAccessDenied             = NewCustomError(10008, "Access denied")                                 // 未授权
	ErrLimitExceed              = NewCustomError(10009, "Resource exhausted")                            // 资源耗尽
	ErrMethodNotAllowed         = NewCustomError(10010, "Method not allowed")                            // 方法不允许
	ErrRequestEntityTooLarge    = NewCustomError(10011, "Request body length over limit")                // 请求体长度超过限制
	ErrJobExpired               = NewCustomError(10012, "Job expired")                                   // 任务超时
	ErrSignParam                = NewCustomError(10013, "Invalid sign")                                  // 签名错误
	ErrValidation               = NewCustomError(10014, "Validation failed")                             // 校验错误
	ErrDatabase                 = NewCustomError(10015, "Database error")                                // 数据库错误
	ErrToken                    = NewCustomError(10016, "Gen token error")                               // 生成 token 错误
	ErrInvalidToken             = NewCustomError(10017, "Invalid token")                                 // 无效的 token
	ErrTokenTimeout             = NewCustomError(10018, "Token timeout")                                 // token 过期
	ErrTooManyRequests          = NewCustomError(10019, "User requests out of rate limit")               // 太多请求
	ErrInvalidTransaction       = NewCustomError(10020, "Invalid transaction")                           // 无效的事务
	ErrEncrypt                  = NewCustomError(10021, "Encrypting the user password error")            // 加密错误
	ErrIPLimit                  = NewCustomError(10022, "IP limit")                                      // IP 限制不能请求该资源
	ErrRequestNotAllow          = NewCustomError(10023, "HTTP method is not supported for this request") // 请求的 HTTP METHOD 不支持，请检查是否选择了正确的 POST/GET 方式
	ErrIPRequestsOutOfRateLimit = NewCustomError(10024, "IP requests out of rate limit")                 // IP 请求频次超过上限
	ErrMissRequiredParam        = NewCustomError(10025, "Miss required parameter (%s)")                  // 缺少必选参数 (%s)，请参考 API 文档
	ErrParamError               = NewCustomError(10026, "Param error, see doc for more info")            // 参数错误，参考 API 文档
	ErrCustomError              = NewCustomError(10027, "Custom error")                                  // 自定义错误
)
