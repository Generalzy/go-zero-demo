package errorx

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func NewApiError(code int, msg string) *ApiError {
	return &ApiError{
		Code: code,
		Msg:  msg,
	}
}

// Error 实现error接口
func (e *ApiError) Error() string {
	return e.Msg
}

// ParamsError 110代表用户系统 1101代表错误码 微服务的错误可以体现系统
var ParamsError = NewApiError(1101, "参数错误")

type ApiErrorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Data ApiErrorResponse打印详细的信息 否则.Error()会被调用只打印一个msg
func (e *ApiError) Data() *ApiErrorResponse {
	return &ApiErrorResponse{
		Code: e.Code,
		Msg:  e.Msg,
	}
}
