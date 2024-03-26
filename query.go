package aliPay

import (
	"encoding/json"
	"github.com/go-tron/local-time"
	"github.com/go-tron/types/fieldUtil"
)

type QueryReq struct {
	TransactionId string `json:"transactionId" validate:"required"`
}

func (u *AliPay) Query(data *QueryReq) (map[string]interface{}, error) {
	if fieldUtil.IsEmpty(data.TransactionId) {
		return nil, ErrorParam("订单号")
	}

	content := map[string]interface{}{
		"out_trade_no": data.TransactionId,
	}
	contentStr, _ := json.Marshal(content)

	params := map[string]interface{}{
		"app_id":      u.AppId,
		"charset":     "utf-8",
		"format":      "JSON",
		"method":      "alipay.trade.query",
		"version":     "1.0",
		"timestamp":   localTime.Now().String(),
		"sign_type":   "RSA2",
		"biz_content": string(contentStr),
	}
	return u.Execute("Query", params)
}
