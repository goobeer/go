package commonhtmlhelper

import (
	"encoding/hex"
	"fasthttpweb/common"
	"fmt"
)

func GenerateVerifyImgLink(verifyUrl, keyVal string) string {
	data, err := common.TripleDesEncrypt([]byte(keyVal), nil)
	if err != nil {
		panic(err.Error())
	}
	return fmt.Sprintf("%s?%s", verifyUrl, hex.EncodeToString(data))
}
