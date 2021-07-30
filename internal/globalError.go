package internal

import "day7/internal/model"

const (
	NoName    = "1001" //用户名为空
	NoAdress  = "1002" //没有websocket地址
	ServerErr = "1003" //内部服务错误
)

// GlobalError 全局异常结构体
type GlobalError struct {
	Code    string `json:"code"`
	Message string
}

//获取err的内容
func (err GlobalError) Error() string {
	return err.Message
}

func NoNameError(message string) GlobalError {
	return GlobalError{
		Code:    NoName,
		Message: message,
	}
}

func NoAdressError(message string) GlobalError {
	return GlobalError{
		Code:    NoAdress,
		Message: message,
	}
}

func ServerError(message string) GlobalError {
	return GlobalError{
		Code:    ServerErr,
		Message: message,
	}
}

func ReturnErr(globalError GlobalError) {
	model.Right.SetText(model.Right.Text + "\n" + globalError.Code + "" + globalError.Message)
}
