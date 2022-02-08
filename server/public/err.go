package public

const (
	ERR_BAD_DATA            = 1001 //数据错误
	ERR_SERVER_PROCESS_MISS = 1002

	ERR_USERNAME_REPEAT    = 2001 //用户名重复
	ERR_USERNAME_NOT_EXIST = 2002
	ERR_PASSWORD_NOT_MATCH = 2003
	ERR_USER_NOT_EXIST     = 2004

	ERR_TOKEN_EMPTY      = 4001
	ERR_TOKEN_EXPIRED    = 4002
	ERR_TOKEN_NOT_HANDLE = 4003
)

var ErrMsg map[uint]string

func init() {
	ErrMsg = make(map[uint]string)
	ErrMsg[ERR_BAD_DATA] = "请求数据错误"
	ErrMsg[ERR_SERVER_PROCESS_MISS] = "服务器内部处理出错"
	ErrMsg[ERR_USERNAME_REPEAT] = "用户名已被占用，请重新选择"
	ErrMsg[ERR_USERNAME_NOT_EXIST] = "用户名不存在"
	ErrMsg[ERR_PASSWORD_NOT_MATCH] = "用户密码错误"
	ErrMsg[ERR_USER_NOT_EXIST] = "用户不存在"
	ErrMsg[ERR_TOKEN_EMPTY] = "请求未携带token,无访问权限"
	ErrMsg[ERR_TOKEN_EXPIRED] = "请求token已过期"
	ErrMsg[ERR_TOKEN_NOT_HANDLE] = "解析token出错"
}
