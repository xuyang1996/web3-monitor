package util

const (
	Unknown        = -1
	Success        = 0
	FailToAccessDB = 1000
	FieldNotSet    = 1001

	FailToGetSignatureOrAccountFromHeader = 2001
	FailToVerifySignature                 = 2002
	FailToGetAccount                      = 2003
	FailToGetIdOrAccountFromPath          = 2004
	InvalidIdArrayForThisAccount          = 2005
	DuplicateUserTasks                    = 2006
)

var MessageMap = map[int]string{
	Unknown:                               "unknown error",
	Success:                               "success",
	FailToAccessDB:                        "fail to access database",
	FailToGetSignatureOrAccountFromHeader: "fail to get signature or account from header",
	FailToVerifySignature:                 "fail to verify signature",
	FailToGetAccount:                      "fail to get account",
	FailToGetIdOrAccountFromPath:          "fail to get id or account from path",
	InvalidIdArrayForThisAccount:          "invalid id array for this account",
	DuplicateUserTasks:                    "duplicate user tasks",
}

const (
	UserActiveStatus       = "ACTIVE"
	UserInactiveStatus     = "INACTIVE"
	UserTaskActiveStatus   = "ACTIVE"
	UserTaskInactiveStatus = "INACTIVE"
	EoaAccount             = "EOA"
	ContractAccount        = "CONTRACT"
)

// BNB 56
// ETH 1
