package aliPay

import (
	localTime "github.com/go-tron/local-time"
	"github.com/go-tron/logger"
	"testing"
)

var api = New(&AliPay{
	Env:   "production",
	AppId: "2021003169604266",
	PrivateKeyPem: `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAlYhIZwNkPGUC4J0N+fdhB4Ww7HE7/O7JjAWTnZ95LPD8MuK5nEP2EYRQCoDD6vDqLvemFBKrCLZAZsF9yl+foCjRJXk/WT7jrQImrvWWy13LZu5cq6utUA6n67XmoPQXgG0fbczghP44BSQ4f992qon9rj2oKSosb16dqfKZmf4MyTiikSxOu3jqUxvfzNf7oR0ah/TgFgZY2nrjccgyvn37EXUhENTn3Gxf9vQt3o1DmnRKwFxzoyqgqHDWN+ZmSFWJzTK8pe6/Yk8b3Lzl4dwUKyKf+k1Q2zP5jqEP12s2lAK1J0QYZHGwkcz1fdekEa2+p6slwRyzdHkqcuIrXwIDAQABAoIBAHGYrcAsUGqdJhly1ppN9xVa1/RXdYYJ9Wz9E45MBydAD6esm+r9qiLWjGPePHfv+0gg9LcNE4ezxKsLVT93c3GdcH1yZdCruRTGrJJ/mcX3BD222QnFiw1lhOXJM2KU7IGw5I5qdSozYmVthcqG/cRCvkgvKN/U3RriGw5vcYcM0lrYeRBOAkBo8MH3S4ATiXU9EkQ2deim8XrC0kLoPfaUlL/xXYZAYfn2En7sdFEavkJGcjarUkQfstGTDXVXDZKJsA7oRY6wK2ydO0ganx6AqwSjqmGoJy6Da42ERggUbldoGqvoIm2OI6RzEX9sgQ6Z/YdlF9A26CwYx3bXtGECgYEAyB837iWR6OGumh3j64DYq/WBUiUiN/lNMAL+dAS18M/+zlthJ8eEhwXHptu1NCquEo2FP09fb8PG1uBiDNyXrBqLPkuTcSK5hrDJIOmV6cJEfx3fOkSO/xAiN7r77bdKnx9V3mYnxJWLwgGMoQHhsYnapvAML7DUWbWqI2OQ7fECgYEAv0jnZuRPF4BbLOIr4PIrBRFOuo2HHci7JMVYCR0kmXFPJcb6l0dMgAcInh3ZGk34IhPDhBXUdMQXc6iHPElYHzl7SAhRYJbjKpcfF19eir+472IZ5utQbsUDpbyhes+uCcGIt3NLSBIhfAJUwQG449L9G0IQQxEzBm0IN9Uunk8CgYBKxp93dMJYajt7ir+nN7W+SzXPI+DtWVHmJrg9UaVKHe3v3WUoH/z9FsPLLT1ACNKSTB8F0PqwIE8j6yO3+pUR0blFxaeKFpeMJHKCwcUqW1SMyvSmKQfldnnSqSOJZ3uSiXrkZvdlFRvrmfiaEMHsPL5eskNbbo9qFd9E6ec1gQKBgBvjcjFzKgDgKurhUspqJFGJ03OpfMCf6oES8KHriNGCTqrQVurFb2bfH6eF7IhEQ+AcB45zbFVV3aF0ObtVai6rP8khxVOSzC4CeHr84ZjTGRB1uhcLyd9MhBqe5OA19Ubg26D7g0dPtWgSIu885Ar7UQGvYRWWJV1TejZBs5lLAoGATLTlacsKUSI4/eH9lKAHqI+FIBMrovBeNaFKrHzHF6N++nrxOLocbu/ygcYNyg2F67AONMFLMb8/jamlPDRx6+9b5ZunCNMvSwFB4/7qLoZ3+Yt+8fGKT68zE8kHfSo0xyOtMNwoP/D8Bqcm0YShU6AjyLDhsu5H69ATKGesIds=
-----END RSA PRIVATE KEY-----`,
	ReturnUrl: "https://alipay.mall.tjgdznzf.com/return",
	NotifyUrl: "https://alipay.mall.tjgdznzf.com/notice",
	Logger:    logger.NewZap("aliPay", "info"),
})

func TestWeb(t *testing.T) {
	result, err := api.Web(&WebReq{
		TransactionId: "20221223010",
		TxnAmount:     0.05,
		TxnTime:       localTime.Now().Ptr(),
		Description:   "测试",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestRefund(t *testing.T) {
	result, err := api.Refund(&RefundReq{
		TransactionId:     "1607281496291291137",
		TxnAmount:         0.06,
		TxnTime:           localTime.Now().Ptr(),
		OrigTransactionId: "1607281496291291136",
		Description:       "撤销",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestQuery(t *testing.T) {
	result, err := api.Query(&QueryReq{
		TransactionId: "1607281496291291136",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestRefundQuery(t *testing.T) {
	result, err := api.RefundQuery(&RefundQueryReq{
		TransactionId:     "1611368660197179392",
		OrigTransactionId: "1611348709520441344",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}
