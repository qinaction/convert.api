package error_wrapper

import (
	"convert.api/libs/common"
	"net/http"
)

type Code int

const (
	SERVER_ERROR Code = 100000 + iota
	NOT_FOUND
	UNKNOWN_ERROR
	PARAMETER_ERROR
	WHITE_LIST
	LEASE_GET_INFO
	LEASE_FORBID
	LEASE_NOT_AUTH
	LEASE_WHITE
	SERVER_LIMITING_ADD_FAIL
	AGGREGATE_NOT_FOUND
)

var (
	errno         = [...]string{"0", "1"}
	CustomizeCode = map[Code]string{
		SERVER_ERROR:             "系统错误",
		NOT_FOUND:                "404未找到",
		UNKNOWN_ERROR:            "未知",
		PARAMETER_ERROR:          "参数错误",
		WHITE_LIST:               "白名单服务错误",
		LEASE_GET_INFO:           "租户信息错误",
		LEASE_NOT_AUTH:           "无访问权限",
		LEASE_FORBID:             "租户被禁用",
		LEASE_WHITE:              "租户不在白名单中",
		SERVER_LIMITING_ADD_FAIL: "服务限流规则添加失败",
		AGGREGATE_NOT_FOUND:      "服务未找到",
	}
)

func ErrCodeToStr(code Code) string {
	return common.IntToStr(int(code))
}

// 500 错误处理
func ServerError() *ErrorException {
	return NewErrorException(http.StatusInternalServerError, ErrCodeToStr(SERVER_ERROR), "OBJECT", http.StatusText(http.StatusInternalServerError), "")
}

// 404 错误
func NotFound() *ErrorException {
	return NewErrorException(http.StatusBadRequest, ErrCodeToStr(NOT_FOUND), "OBJECT", http.StatusText(http.StatusNotFound), "")
}

// 未知错误
func UnknownError(message string) *ErrorException {
	return NewErrorException(http.StatusForbidden, ErrCodeToStr(UNKNOWN_ERROR), "OBJECT", message, "")
}

// 参数错误
func ParameterError(message string) *ErrorException {
	return NewErrorException(http.StatusBadRequest, ErrCodeToStr(PARAMETER_ERROR), "OBJECT", message, "")
}

//成功时返回对象
func WithSuccessObj(data interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[0], "OBJECT", "", data)
}

//成功时返回数组
func WithSuccess(data ...interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[0], "OBJECT", "", data)
}

//错误时返回
func WitheErrorObj(message string, data interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[1], "OBJECT", message, data)
}

//错误时返回数组
func WitheError(message string, data ...interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[1], "OBJECT", message, data)
}

//白名单没设置
func WhiteListError(message string, data ...interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, ErrCodeToStr(WHITE_LIST), "OBJECT", message, data)
}

//自定义错误码
func ErrorCodeObj(errorCode Code, data interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, ErrCodeToStr(errorCode), "OBJECT", CustomizeCode[errorCode], data)
}

//单个成功时返回 json
func WithSingle(data interface{}) *ErrorException {
	return NewErrorException(http.StatusOK, errno[0], "OBJECT", "", data)
}
