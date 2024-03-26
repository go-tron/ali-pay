package sdkConfig

var Testing = New(&SDKConfig{
	SignPublicKeyPem: `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAiX4wygj4rPO4IzQJWMnevsopDn1jYx1d5vOkUXkb946mU1SmNLALRcD6TLlTVUEQ/JFUVU+5nz47YN8fIO6atvb2ctB884SYUpml0VVdR60aNfCauHU+nFxV5hV7PED8zkbrel67UZfT8AcMAPnVxDsVo3Ng1TeFOopc/V00nV+QWWZIcrp+vNx480GC1OUWnIAk4dDCaePikm3clLmC9ENFaZ2pKeC4BBDkxa8f70RK/WqQJygNwRajND0N4RiFhgE7b6KQdCMPP2IZWsofCP+BMADY2iG50tarbRnGrWtmyoRSk1CNOxrGpq9rXNubugpoKCmtNJ8lboEco64SLQIDAQAB
-----END PUBLIC KEY-----`,
	//请求地址
	GateWay: "https://openapi.alipay.com/gateway.do",
	ResponseProperty: map[string]string{
		"Query":       "alipay_trade_query_response",
		"Refund":      "alipay_trade_refund_response",
		"RefundQuery": "alipay_trade_fastpay_refund_query_response",
	},
})
