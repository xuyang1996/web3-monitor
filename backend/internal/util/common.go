package util

import (
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/storyicon/sigverify"
)

const (
	SIGNATURE_DATA = "hello"
)

func VerifySignature(address string, signature string) bool {
	valid, _ := sigverify.VerifyEllipticCurveHexSignatureEx(
		ethcommon.HexToAddress(address),
		[]byte(SIGNATURE_DATA),
		signature,
	)
	return valid
}
