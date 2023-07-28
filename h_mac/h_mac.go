package h_mac

import (
	"bytes"
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

	var m map[string]interface{}
	if err = json.Unmarshal(body, &m); err != nil {
		return false, err
	}

	// 请求参数排序
	bodyBytes, err := json.Marshal(m)
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

func GenerateHMAC(requestURI, key string, body []byte) (error, string) {
	var tempStruct map[string]interface{}
	if err := json.Unmarshal(body, &tempStruct); err != nil {
		return errors.New("invalid body data"), ""
	}
	newBody, err := json.Marshal(tempStruct)
	if err != nil {
		return errors.New(fmt.Sprintf("marshal body data failed, err: %v", err)), ""
	}
	byteData := []byte(requestURI)
	for _, b := range newBody {
		byteData = append(byteData, b)
	}
	tempByteData := string(byteData)
	byteData, err = json.Marshal(tempByteData)
	if err != nil {
		return errors.New(fmt.Sprintf("marshal temp data failed, err: %v", err)), ""
	}

	return nil, CalculateHMAC(byteData, key)
}
