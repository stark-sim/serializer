package code

type MyCode int64

const (
	ServerBusy MyCode = 10099

	Success        MyCode = 20000
	SuccessCreated MyCode = 20001

	InvalidParams MyCode = 30000
	NotFound      MyCode = 30001
	NotEnough     MyCode = 30002
	SourceExist   MyCode = 30003

	AuthFailed     MyCode = 40000
	UnLogin        MyCode = 40001
	InvalidKey     MyCode = 40002
	InvalidRequest MyCode = 40003

	ServerErr              MyCode = 50000
	ServerErrDB            MyCode = 50001
	ServerErrCache         MyCode = 50002
	ServerErrThirdPartyAPI MyCode = 50003

	FailGetInviteCode MyCode = 50050
	FailHasRegister   MyCode = 50051
)

var msgFlags = map[MyCode]string{
	ServerBusy: "Busy service",

	Success:        "Success",
	SuccessCreated: "Create Success",

	InvalidParams: "Illegal or missing parameters",
	NotFound:      "Resource does not exist",
	NotEnough:     "Insufficient resources to perform the operation",
	SourceExist:   "Resource already exists",

	AuthFailed:     "Login has expired, please log in again",
	UnLogin:        "Not logged in",
	InvalidKey:     "Illegal key",
	InvalidRequest: "Illegal request",

	ServerErr:              "Server exception",
	ServerErrDB:            "Server database exception",
	ServerErrCache:         "Cache exception",
	ServerErrThirdPartyAPI: "Third-party interface call exception",
	FailGetInviteCode:      "Invitation code does not exist",
	FailHasRegister:        "This phone number has been registered",
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
