package aliPay

import (
	"crypto"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/go-tron/ali-pay/sdkConfig"
	baseError "github.com/go-tron/base-error"
	"github.com/go-tron/crypto/encoding"
	"github.com/go-tron/crypto/rsaUtil"
	"github.com/go-tron/logger"
	"github.com/go-tron/types/mapUtil"
	"github.com/tidwall/gjson"
	"net/url"
)

var (
	ErrorParam   = baseError.SystemFactory("9001", "支付参数错误:{}")
	ErrorSign    = baseError.System("9004", "支付签名失败")
	ErrorRequest = baseError.System("9005", "支付服务连接失败")
	ErrorVerify  = baseError.System("9007", "支付验证失败")
	ErrorCode    = baseError.SystemFactory("9010")
)

func New(config *AliPay) *AliPay {
	if config.Env == "" || config.AppId == "" || config.PrivateKeyPem == "" || config.Logger == nil {
		panic("invalid aliPay config")
	}

	privateKey, err := rsaUtil.GetPrivateKeyPem([]byte(config.PrivateKeyPem))
	if err != nil {
		panic(err)
	}
	config.PrivateKey = privateKey

	if config.Env == "production" {
		config.SDKConfig = sdkConfig.Production
	} else {
		config.SDKConfig = sdkConfig.Testing
	}
	return config
}

type AliPay struct {
	Env           string
	AppId         string
	PrivateKeyPem string
	PrivateKey    *rsa.PrivateKey
	SDKConfig     *sdkConfig.SDKConfig
	ReturnUrl     string
	NotifyUrl     string
	Logger        logger.Logger
}

func (u *AliPay) GetUrl(obj map[string]interface{}) string {
	values := url.Values{}
	for k, v := range obj {
		values.Add(k, fmt.Sprint(v))
	}
	return u.SDKConfig.GateWay + "?" + values.Encode()
}

func (u *AliPay) Sign(obj map[string]interface{}) error {
	signStr := mapUtil.ToSortString(obj)
	sign, err := rsaUtil.Sign(signStr, u.PrivateKey, crypto.SHA256, &encoding.Base64{})
	if err != nil {
		return ErrorSign
	}
	obj["sign"] = sign
	return nil
}

func (u *AliPay) VerifyReq(signStr string, sign string) error {
	if err := rsaUtil.Verify(signStr, sign, u.SDKConfig.SignPublicKey, crypto.SHA256, &encoding.Base64{}); err != nil {
		return ErrorVerify
	}
	return nil
}

func (u *AliPay) Verify(obj map[string]interface{}) error {
	sign := obj["sign"]
	if sign == nil {
		return ErrorVerify
	}
	delete(obj, "sign")
	delete(obj, "sign_type")
	signStr := mapUtil.ToSortString(obj)

	if err := rsaUtil.Verify(signStr, sign.(string), u.SDKConfig.SignPublicKey, crypto.SHA256, &encoding.Base64{}); err != nil {
		return ErrorVerify
	}
	return nil
}

func (u *AliPay) Execute(name string, data map[string]interface{}) (m map[string]interface{}, err error) {
	if err := u.Sign(data); err != nil {
		return nil, err
	}

	url := u.GetUrl(data)

	request, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	u.Logger.Info(string(request),
		u.Logger.Field("orderId", data["orderId"]),
		u.Logger.Field("name", name),
		u.Logger.Field("type", "request"),
	)

	if name == "Web" {
		m, err = u.Url(name, url, data)
	} else {
		m, err = u.Request(name, url, data)
	}

	return m, err
}

func (u *AliPay) Request(name string, url string, data map[string]interface{}) (result map[string]interface{}, err error) {
	var (
		response = ""
	)
	defer func() {
		u.Logger.Info(response,
			u.Logger.Field("orderId", data["orderId"]),
			u.Logger.Field("name", name),
			u.Logger.Field("type", "response"),
			u.Logger.Field("error", err))
	}()

	client := resty.New().R()
	resp, err := client.Post(url)
	if err != nil {
		return nil, ErrorRequest
	}

	response = gjson.Get(string(resp.Body()), u.SDKConfig.ResponseProperty[name]).String()

	var res = make(map[string]interface{})
	if err := json.Unmarshal([]byte(response), &res); err != nil {
		return nil, err
	}
	if err := u.VerifyReq(response, gjson.Get(string(resp.Body()), "sign").String()); err != nil {
		return nil, err
	}
	if res["code"] != "10000" {
		return nil, ErrorCode(res["sub_msg"])
	}
	return res, nil
}

func (u *AliPay) Url(name string, url string, data map[string]interface{}) (map[string]interface{}, error) {
	return map[string]interface{}{
		"url": url,
	}, nil
}
