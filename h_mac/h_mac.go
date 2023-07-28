package h_mac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
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

	body, err := c.GetRawData()
	if err != nil {
		return false, err
	}

	var m map[string]interface{}
	if err = json.Unmarshal(body, &m); err != nil {
		return false, err
	}

	// 请求参数排序
	bodyBytes, err := json.Marshal(m)
	if err != nil {
		return false, err
	}

	// 整合需要序列化的数据
	strData := c.Request.RequestURI + string(bodyBytes)
	data, err := json.Marshal(strData)
	if err != nil {
		return false, err
	}

	return receivedHMAC == CalculateHMAC(data, key), nil
}

func GenerateHMAC(requestURI, key string, body []byte) (string, error) {
	var tempStruct map[string]interface{}
	if err := json.Unmarshal(body, &tempStruct); err != nil {
		return "", errors.New("invalid body data")
	}
	newBody, err := json.Marshal(tempStruct)
	if err != nil {
		return "", errors.New(fmt.Sprintf("marshal body data failed, err: %v", err))
	}

	strData := requestURI + string(newBody)
	byteData, err := json.Marshal(strData)
	if err != nil {
		return "", errors.New(fmt.Sprintf("marshal str data failed, err: %v", err))
	}

	return CalculateHMAC(byteData, key), nil
}
