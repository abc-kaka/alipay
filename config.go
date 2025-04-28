package alipay

import "time"

// Config 支付宝配置
type Config struct {
	AppId        string // 支付宝分配给开发者的应用ID
	AppAuthToken string // 授权令牌
	PrivateKey   string // 私钥
	PublicKey    string // 公钥

	ServerUrl string
	Format    string
	Charset   string
	SignType  string
}

func NewConfig() *Config {
	return &Config{
		AppId:        "",
		AppAuthToken: "",
		PrivateKey:   "",
		PublicKey:    "",
		ServerUrl:    "https://openapi.alipay.com/gateway.do",
		Format:       "json",
		Charset:      "UTF-8",
		SignType:     "RSA2",
	}
}

// GetApiPublicParams 获取api公共参数
func (config *Config) GetApiPublicParams(apiCode string) (paramsMap map[string]string) {
	paramsMap = make(map[string]string)
	paramsMap["app_id"] = config.AppId
	paramsMap["app_auth_token"] = config.AppAuthToken
	paramsMap["method"] = apiCode
	paramsMap["format"] = config.Format
	paramsMap["charset"] = config.Charset
	paramsMap["sign_type"] = config.SignType
	paramsMap["timestamp"] = time.Now().Format("2006-01-02 15:04:05")
	paramsMap["version"] = "1.0"
	return
}
