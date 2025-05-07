package util

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
)

func VerifyNotify(rawForm string, publicKey string) (params map[string]any, verify bool, err error) {
	// 参数集
	values, _ := url.ParseQuery(rawForm)
	err = verifySign(values, publicKey)
	if err != nil {
		return
	}
	params = make(map[string]any)
	for k, v := range values {
		if len(v) > 0 {
			params[k] = v[0]
		}
	}
	verify = true
	return
}

func verifySign(params url.Values, alipayPublicKey string) error {
	sign := params.Get("sign")
	signType := params.Get("sign_type")

	if sign == "" || signType != "RSA2" {
		return errors.New("missing sign or unsupported sign_type")
	}

	// 1. 排除 sign 和 sign_type 字段
	keys := make([]string, 0, len(params))
	for k := range params {
		if k != "sign" && k != "sign_type" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	// 2. 构造待签名字符串
	buf := bytes.NewBuffer(make([]byte, 0, len(keys)))
	for i, k := range keys {
		// 拼接签名字符串
		buf.WriteString(k + "=" + params.Get(k))
		if i < len(keys)-1 {
			buf.WriteString("&")
		}
	}
	signStr, err := url.QueryUnescape(buf.String())
	if err != nil {
		return err
	}

	// 3. 解码签名
	sigBytes, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return fmt.Errorf("decode signature error: %w", err)
	}
	// 4. 解析支付宝公钥
	block, _ := pem.Decode([]byte(alipayPublicKey))
	rsaPubKeyAny, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("解析公钥失败！")
	}
	rsaPubKey := rsaPubKeyAny.(*rsa.PublicKey)

	// 5. 做 SHA256 摘要
	hash := sha256.New()
	hash.Write([]byte(signStr))
	hashed := hash.Sum(nil)

	// 6. 验签
	err = rsa.VerifyPKCS1v15(rsaPubKey, crypto.SHA256, hashed, sigBytes)
	if err != nil {
		return fmt.Errorf("rsa verify error: %w", err)
	}
	return nil
}

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

// 自动尝试 GBK → UTF-8（如果原始就是 UTF-8，会保持不变）
func EnsureUTF8(body []byte) []byte {
	// 尝试解码成 UTF-8，如果失败说明原始就是 UTF-8
	reader := transform.NewReader(bytes.NewReader(body), simplifiedchinese.GBK.NewDecoder())
	result, err := ioutil.ReadAll(reader)
	if err != nil {
		return body // 说明原始就是 UTF-8，无需转换
	}
	fmt.Println("转换前：", string(body))
	fmt.Println("转换后：", string(result))
	return result
}
