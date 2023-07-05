package code_msgs

import "fmt"

type CodeMsg struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// Success 200** 成功系列
var (
	Success        = CodeMsg{Code: 20000, Msg: "成功"}
	SuccessCreated = CodeMsg{Code: 20001, Msg: "创建成功"}
)

// InputFail 300** 用户原因错误系列
var (
	InputFail     = CodeMsg{Code: 30000, Msg: "输入参数错误"}
	UnmarshalFail = CodeMsg{30001, "反序列化失败"}
	NotEnough     = CodeMsg{30002, "数量不足"}
)

// ValidFail 400** 权限错误系列
var (
	ValidFail          = CodeMsg{Code: 40000, Msg: "权限错误"}
	AccessTokenExpire  = CodeMsg{40001, "access_token 过期"}
	RefreshTokenExpire = CodeMsg{40002, "refresh_token 过期"}
	UnLogin_           = CodeMsg{40003, "未登录"}
	KeyInvalid         = CodeMsg{40004, "key 无效"}
)

// Fail 500** 内部错误系列
var (
	Fail                    = CodeMsg{50000, "失败"}
	FailDB                  = CodeMsg{50001, "数据库处理异常"}
	FailThirdPartyAPI       = CodeMsg{50002, "第三方接口异常"}
	FailLogin               = CodeMsg{50003, "登陆失败"}
	FailCreate              = CodeMsg{50004, "创建失败"}
	FailGetOrder            = CodeMsg{50005, "获取订单失败"}
	FailCallBack            = CodeMsg{50006, "回调出错"}
	FailQuery               = CodeMsg{50007, "查询出错"}
	FailSetUserInfo         = CodeMsg{50008, "设置用户信息出错"}
	FailTranslate           = CodeMsg{50009, "翻译失败"}
	FailCollect             = CodeMsg{50010, "收藏失败"}
	FailListCollect         = CodeMsg{50011, "获取收藏列表失败"}
	FailDeleteCollect       = CodeMsg{50012, "删除收藏项失败"}
	FailEncrypt             = CodeMsg{50013, "加密失败"}
	FailMakeMissionOrder    = CodeMsg{50014, "创建任务订单失败"}
	FailConfirmMissionOrder = CodeMsg{50015, "确认任务订单状态失败"}
	FailTranslateMachine    = CodeMsg{50016, "关联机器设备失败"}
	FailMakeProfitBill      = CodeMsg{50017, "创建分润账单失败"}
	FailGetProfitBill       = CodeMsg{50018, "获取分润账单失败"}
	FailGetHmacKey          = CodeMsg{50019, "获取密钥对失败"}
)

func (e *CodeMsg) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Msg)
}
