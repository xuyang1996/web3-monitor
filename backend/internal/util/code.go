package util

const (
	Unknown                     = -1
	Success                     = 0
	FailToGetSignatureOrAccount = 1001
)

var MessageMap = map[int]string{
	Unknown:                     "unknown error",
	Success:                     "success",
	FailToGetSignatureOrAccount: "fail to get signature or account from header",
}
