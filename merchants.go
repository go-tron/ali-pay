package aliPay

type IMerchants interface {
	GetMerchantById(string) (*AliPay, error)
}

type Merchants struct {
	Merchants IMerchants
}

func (u *Merchants) Web(merId string, data *WebReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Web(data)
}

func (u *Merchants) Verify(merId string, data map[string]interface{}) error {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return err
	}
	return merchant.Verify(data)
}

func (u *Merchants) Query(merId string, data *QueryReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Query(data)
}

func (u *Merchants) Refund(merId string, data *RefundReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.Refund(data)
}

func (u *Merchants) RefundQuery(merId string, data *RefundQueryReq) (map[string]interface{}, error) {
	merchant, err := u.Merchants.GetMerchantById(merId)
	if err != nil {
		return nil, err
	}
	return merchant.RefundQuery(data)
}
