package aliPay

import (
	"encoding/json"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
)

type RefundReq struct {
	TransactionId     string          `json:"transactionId" validate:"required"`
	TxnAmount         float64         `json:"txnAmount" validate:"required"`
	TxnTime           *localTime.Time `json:"txnTime" validate:"required"`
	OrigTransactionId string          `json:"origTransactionId" validate:"required"`
	Description       string          `json:"description" validate:"required"`
}

func (u *AliPay) Refund(data *RefundReq) (map[string]interface{}, error) {

	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}
	if fieldUtil.IsEmpty(data.TxnAmount) {
		return nil, ErrorParam("支付金额")
	}
	if fieldUtil.IsEmpty(data.TxnTime) {
		return nil, ErrorParam("退款时间")
	}
	if fieldUtil.IsEmpty(data.OrigTransactionId) {
		return nil, ErrorParam("原交易单号")
	}
	if fieldUtil.IsEmpty(data.Description) {
		return nil, ErrorParam("订单详情")
	}

	content := map[string]interface{}{
		"out_trade_no":   data.OrigTransactionId,
		"out_request_no": data.TransactionId,
		"refund_amount":  data.TxnAmount,
		"refund_reason":  data.Description,
	}
	contentStr, _ := json.Marshal(content)

	params := map[string]interface{}{
		"app_id":      u.AppId,
		"charset":     "utf-8",
		"format":      "JSON",
		"method":      "alipay.trade.refund",
		"version":     "1.0",
		"timestamp":   data.TxnTime.String(),
		"sign_type":   "RSA2",
		"biz_content": string(contentStr),
	}
	return u.Execute("Refund", params)
}
