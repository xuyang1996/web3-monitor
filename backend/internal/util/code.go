package util

const (
	Unknown                     = -1
	Success                     = 0
	FailToAccessDB              = 1000
	FieldNotSet                 = 2000
	FailToGetSignatureOrAccount = 2001
	FailToVerifySignature       = 2002
)

var MessageMap = map[int]string{
	Unknown:                     "unknown error",
	Success:                     "success",
	FailToAccessDB:              "fail to access database",
	FailToGetSignatureOrAccount: "fail to get signature or account from header",
	FailToVerifySignature:       "fail to verify signature",
}

const (
	UserActiveStatus       = "active"
	UserInactiveStatus     = "inactive"
	UserTaskActiveStatus   = "active"
	UserTaskInactiveStatus = "inactive"
)
