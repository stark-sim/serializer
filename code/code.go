package code

type MyCode int64

const (
	ServerBusy MyCode = 10099

	Success        MyCode = 20000
	SuccessCreated MyCode = 20001

	InvalidParams MyCode = 30000
	NotFound      MyCode = 30001
	NotEnough     MyCode = 30002

	AuthFailed MyCode = 40000
	UnLogin    MyCode = 40001
	InvalidKey MyCode = 40002

	ServerErr              MyCode = 50000
	ServerErrDB            MyCode = 50001
	ServerErrCache         MyCode = 50002
	ServerErrThirdPartyAPI MyCode = 50003

	FailGetInviteCode MyCode = 50050
	FailHasRegister   MyCode = 50051
)

var msgFlags = map[MyCode]string{
	ServerBusy: "服务繁忙",

	Success:        "成功",
	SuccessCreated: "创建成功",

	InvalidParams: "非法参数或缺失",
	NotFound:      "资源不存在",
	NotEnough:     "资源不足以执行该操作",

	AuthFailed: "权限验证失败",
	UnLogin:    "未登录",
	InvalidKey: "非法秘钥",

	ServerErr:              "服务端异常",
	ServerErrDB:            "服务端数据库异常",
	ServerErrCache:         "缓存异常",
	ServerErrThirdPartyAPI: "第三方接口调用异常",
	FailGetInviteCode:      "邀请码不存在",
	FailHasRegister:        "该手机号已注册",
}

func (c MyCode) Msg() string {
	msg, ok := msgFlags[c]
	if ok {
		return msg
	}
	return msgFlags[ServerBusy]
}

func (c MyCode) Error() string {
	return c.Msg()
}
