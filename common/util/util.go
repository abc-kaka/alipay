package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

// GenerateSign 生成签名1
func GenerateSign(paramMap map[string]string, privateKeyStr string) (signStr string, err error) {
	// 参数转字符
	paramMapStr := MapToStr(paramMap)
	// 解析私钥
	privateBlock, _ := pem.Decode([]byte(privateKeyStr))
	privateKey, err := x509.ParsePKCS1PrivateKey(privateBlock.Bytes)
	if err != nil {
		privateKey8, err8 := x509.ParsePKCS8PrivateKey(privateBlock.Bytes)
		if err8 != nil {
			err = fmt.Errorf("解析私钥失败！")
			return
		}
		privateKey = privateKey8.(*rsa.PrivateKey)
	}
	// RSA2Sign 生成RSA2签名
	h := sha256.New()
	h.Write([]byte(paramMapStr))
	digest := h.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
	if err != nil {
		err = fmt.Errorf("生成RSA2签名失败！")
		return
	}
	signStr = base64.StdEncoding.EncodeToString(sign)
	return
}

// HttpRequest Http请求
func HttpRequest(method string, path string, paramMap map[string]string) (bodyStr string, err error) {
	// 创建一个URL对象
	u, err := url.Parse(path)
	if err != nil {
		return
	}
	// 添加查询参数
	params := url.Values{}
	for key, value := range paramMap {
		params.Add(key, value)
	}
	u.RawQuery = params.Encode()
	// 构建请求
	req, err := http.NewRequest(method, u.String(), nil)
	if err != nil {
		return
	}
	// 发起请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	// 读响应
	body, err := ioutil.ReadAll(resp.Body)
	bodyStr = string(body)
	return
}

// MapToStr map转字符
func MapToStr(dataMap map[string]string) (mapStr string) {
	keys := make([]string, 0, len(dataMap))
	for k := range dataMap {
		keys = append(keys, k)
	}
	// 对切片进行排序
	sort.Strings(keys)
	// 根据排序后的键遍历map并打印值
	for _, k := range keys {
		mapStr += k + "=" + dataMap[k] + "&"
	}
	mapStr = mapStr[:len(mapStr)-1]
	return
}

func MapToStruct(data map[string]interface{}, res interface{}) (err error) {
	// 正常数据封装
	dataByte, err := json.Marshal(data)
	if err != nil {
		return
	}
	err = json.Unmarshal(dataByte, &res)
	return
}
