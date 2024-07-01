package sdk

import (
	"context"
	"fmt"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/services/payments/jsapi"
	"github.com/wechatpay-apiv3/wechatpay-go/services/refunddomestic"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
	"log"
	"wechat-pay/io"
	"wechat-pay/pojo"
)

var (
	mchID                      string // 商户号
	mchCertificateSerialNumber string // 商户证书序列号
	mchAPIv3Key                string // 商户APIv3密钥
	appId                      string
	notifyUrl                  string
	ctx                        = context.Background()
	api                        jsapi.JsapiApiService
	refs                       refunddomestic.RefundsApiService
)

func init() {

	initConf()

	mchPrivateKey, err := utils.LoadPrivateKeyWithPath("apiclient_key.pem")
	if err != nil {
		log.Fatal("load merchant private key error")
	}

	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mchID, mchCertificateSerialNumber, mchPrivateKey, mchAPIv3Key),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		log.Fatalf("new wechat sdk client err:%s", err)
	}

	svc := jsapi.JsapiApiService{Client: client}
	ref := refunddomestic.RefundsApiService{Client: client}
	api = svc
	refs = ref
}

func initConf() {
	conf, err := io.ReadFile()
	if err != nil {
		panic(err)
	}
	mchID = conf["mchID"].(string)
	mchCertificateSerialNumber = conf["mchCertificateSerialNumber"].(string)
	mchAPIv3Key = conf["mchAPIv3Key"].(string)
	appId = conf["appId"].(string)
	notifyUrl = conf["notifyUrl"].(string)
}

func Payment(body pojo.Pay) any {
	resp, _, err := api.PrepayWithRequestPayment(ctx,
		jsapi.PrepayRequest{
			Appid:       core.String(appId),
			Mchid:       core.String(mchID),
			Description: core.String(body.Desc),
			OutTradeNo:  core.String(body.OrderId),
			NotifyUrl:   core.String(notifyUrl),
			Amount: &jsapi.Amount{
				Total: core.Int64(body.Price),
			},
			Payer: &jsapi.Payer{
				Openid: core.String(body.OpenId),
			},
		},
	)

	if err == nil {
		return resp
	}

	return err
}

func Status(orderId string) any {

	resp, _, err := api.QueryOrderByOutTradeNo(ctx,
		jsapi.QueryOrderByOutTradeNoRequest{
			OutTradeNo: core.String(orderId),
			Mchid:      core.String(mchID),
		},
	)

	if err != nil {
		fmt.Println(resp)
		return err
	}
	fmt.Println(resp)
	return resp
}

func Close(orderId string) any {
	result, err := api.CloseOrder(ctx,
		jsapi.CloseOrderRequest{
			OutTradeNo: core.String(orderId),
			Mchid:      core.String(mchID),
		},
	)

	if err != nil {
		return err
	}
	return result.Response.StatusCode
}

func Refund(body pojo.Refund) any {
	_, result, err := refs.Create(ctx,
		refunddomestic.CreateRequest{
			OutTradeNo:  core.String(body.OrderId),
			OutRefundNo: core.String(body.OrderId),
			Amount: &refunddomestic.AmountReq{
				Currency: core.String("CNY"),
				Refund:   core.Int64(body.RefundAmount),
				Total:    core.Int64(body.TotalAmount),
			},
		},
	)
	if err != nil {
		return err
	}
	return result.Response.StatusCode
}
