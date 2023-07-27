package h_mac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	HeaderNameHMAC = "HMAC"
)

func calculateHMAC(data []byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC 验证 HMAC
func VerifyHMAC(c *gin.Context, key string) (bool, error) {
	receivedHMAC := c.GetHeader(HeaderNameHMAC)
	// 从请求中读取原始数据
	fmt.Printf("URI is %v\n", c.Request.RequestURI)
	//input := []byte(uri)
	body, err := c.Copy().GetRawData()
	if err != nil {
		return false, err
	}

	return receivedHMAC == calculateHMAC(body, key), nil
}
