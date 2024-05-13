package resp

type MyCode int64 //自定义MyCode类型

const ( //使用 const 关键字定义了一组错误码，每个错误码都是 MyCode 类型的常量。
	CodeSuccess         MyCode = 1000
	CodeInvalidParams   MyCode = 1001
	CodeUserExist       MyCode = 1002
	CodeUserNotExist    MyCode = 1003
	CodeInvalidPassword MyCode = 1004
	CodeServerBusy      MyCode = 1005

	CodeInvalidToken      MyCode = 1006
	CodeInvalidAuthFormat MyCode = 1007
	CodeNotLogin          MyCode = 1008
	CodeNeedLogin         MyCode = 1009
)

var msgFlags = map[MyCode]string{ //map类型,映射错误码和信息
	CodeSuccess:         "success",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",

	CodeInvalidToken:      "无效的Token",
	CodeInvalidAuthFormat: "认证格式有误",
	CodeNotLogin:          "未登录",
	CodeNeedLogin:         "需要登录",
}

func (c MyCode) Msg() string { //把上边两个连起来，传入code，传出code对应的信息
	//当调用 Msg 方法时，会查找 msgFlags 中是否存在对应的错误信息，
	//如果存在则返回，否则返回默认的错误信息，例如 CodeServerBusy 对应的信息。
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[CodeServerBusy]
}
