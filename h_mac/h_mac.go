package h_mac

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

const (
	HeaderNameHMAC = "HMAC"
)

func CalculateHMAC(data []byte, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

// VerifyHMAC 验证 HMAC
func VerifyHMAC(c *gin.Context, key string) (bool, error) {
	receivedHMAC := c.GetHeader(HeaderNameHMAC)
	// 从请求中读取原始数据
	fmt.Printf("URI is %v\n", c.Request.RequestURI)
	uriBytes, err := json.Marshal(c.Request.RequestURI)
	if err != nil {
		return false, err
	}

	body, err := c.Copy().GetRawData()
	if err != nil {
		return false, err
	}

	// 请求参数排序
	bodyBytes, err := json.Marshal(string(body))
	if err != nil {
		return false, err
	}

	// 合并数据
	byteList := make([][]byte, 2)
	byteList[0] = uriBytes
	byteList[1] = bodyBytes
	data := bytes.Join(byteList, []byte{})

	return receivedHMAC == CalculateHMAC(data, key), nil
}
