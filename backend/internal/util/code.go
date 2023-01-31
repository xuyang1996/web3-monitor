package util

const (
	Unknown                     = -1
	Success                     = 0
	FieldNotSet                 = 1000
	FailToGetSignatureOrAccount = 1001
	FailToVerifySignature       = 1002
)

var MessageMap = map[int]string{
	Unknown:                     "unknown error",
	Success:                     "success",
	FailToGetSignatureOrAccount: "fail to get signature or account from header",
	FailToVerifySignature:       "fail to verify signature",
}

const (
	UserActiveStatus       = "active"
	UserInactiveStatus     = "inactive"
	UserTaskActiveStatus   = "active"
	UserTaskInactiveStatus = "inactive"
)
