package main

import (
	"fmt"

	"github.com/storyicon/sigverify"
)

func main() {
	// valid, err := sigverify.VerifyEllipticCurveHexSignatureEx(
	// 	ethcommon.HexToAddress("0x8Cf1a46E6d90Dfe3471970e057235Df619666666"),
	// 	[]byte("hello"),
	// 	"0x253e557270a41ca3a6cc121bc030974cb82a41108712f66d0309707c27a5bf834aecb9f5dd805e277f3e038505ffae931623f5ba3dc951cf4daac1b369be11571c",
	// )
	// fmt.Println(valid, err) // true <nil>

	sigBytes, err := sigverify.HexDecode("0x253e557270a41ca3a6cc121bc030974cb82a41108712f66d0309707c27a5bf834aecb9f5dd805e277f3e038505ffae931623f5ba3dc951cf4daac1b369be11571c")
	if err != nil {
		fmt.Println(err)
	}
	address, err := sigverify.EcRecover([]byte("hello"), sigBytes)
	fmt.Println(address, err)
}
