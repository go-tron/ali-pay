package aliPay

import (
	"encoding/json"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
)

type WebReq struct {
	TransactionId string          `json:"transactionId" validate:"required"`
	TxnAmount     float64         `json:"txnAmount" validate:"required"`
	TxnTime       *localTime.Time `json:"txnTime" validate:"required"`
	Description   string          `json:"description" validate:"required"`
	ReturnUrl     string          `json:"returnUrl"`
	NotifyUrl     string          `json:"notifyUrl"`
}

func (u *AliPay) Web(data *WebReq) (map[string]interface{}, error) {
	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.TxnTime) {
		return nil, ErrorParam("支付时间")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	content := map[string]interface{}{
		"body":         data.Description,
		"subject":      data.Description,
		"out_trade_no": data.TransactionId,
		"total_amount": data.TxnAmount,
		"product_code": "QUICK_WAP_PAY",
	}
	contentStr, _ := json.Marshal(content)

	params := map[string]interface{}{
		"app_id":      u.AppId,
		"charset":     "utf-8",
		"format":      "JSON",
		"method":      "alipay.trade.wap.pay",
		"version":     "1.0",
		"timestamp":   data.TxnTime.String(),
		"sign_type":   "RSA2",
		"return_url":  u.ReturnUrl,
		"notify_url":  u.NotifyUrl,
		"biz_content": string(contentStr),
	}
	if data.ReturnUrl != "" {
		params["return_url"] = data.ReturnUrl
	}
	if data.NotifyUrl != "" {
		params["notify_url"] = data.NotifyUrl
	}
	return u.Execute("Web", params)
}
