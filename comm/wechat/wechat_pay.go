package wechat

import (
	"math/rand"
	"strconv"
	"time"
	//	"github.com/chanxuehong/utils"
	wxcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	wxpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

type WxPayServer struct {
	WxAppId   string
	MchId     string
	ApiKey    string
	NotifyUrl string
}

type WxResponse struct {
	Appid         string `xml:"appid"`
	BankType      string `xml:"bank_type"`
	CashFee       int    `xml:"cash_fee"`
	FeeType       string `xml:"fee_type"`
	IsSubscribe   string `xml:"is_subscribe"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"NonceStr"`
	OpenID        string `xml:"openid"`
	OutTradeNo    string `xml:"out_trade_no"`
	ResultCode    string `xml:"result_code"`
	ReturnCode    string `xml:"return_code"`
	Sign          string `xml:"sign"`
	TimeEnd       string `xml:"time_end"`
	TotalFee      int    `xml:"total_fee"`
	TradeType     string `xml:"trade_type"`
	TransactionId string `xml:"transaction_id"`
}

var wxPayServer *WxPayServer

func init() {
	initservice()
}

func initservice() {
	if wxPayServer == nil {
		wxPayServer = new(WxPayServer)
	}
	wxPayServer.WxAppId = "wxa2b17d163cc88e9d"
	wxPayServer.MchId = "1496488452"
	wxPayServer.ApiKey = "asdfghjklpoiuytrewqzxcvbnm123456"
	wxPayServer.NotifyUrl = "http://testing.foxhelper.cn/api/wechat/pay/callback"

}

//获取微信服务
func GetwxPayService() *WxPayServer {
	if wxPayServer == nil {
		initservice()
	}

	return wxPayServer
}

func (serv *WxPayServer) UnifiedOrder(money int, name, clientIp, openid, fee_type, trade_type, notifyurl, outTradeNo string) (map[string]string, error) {
	ret := make(map[string]string)
	client := wxcore.NewClient(serv.WxAppId, serv.MchId, serv.ApiKey, nil)
	var order wxpay.UnifiedOrderRequest
	order.Body = name
	order.FeeType = fee_type
	order.NotifyURL = notifyurl
	order.TotalFee = int64(money)
	order.SpbillCreateIP = clientIp
	order.TradeType = trade_type
	order.OpenId = openid
	order.OutTradeNo = outTradeNo
	if resp, err := wxpay.UnifiedOrder2(client, &order); err != nil {
		return ret, err
	} else {
		ret["appId"] = serv.WxAppId
		ret["timeStamp"] = strconv.FormatInt(time.Now().Unix(), 10)
		ret["nonceStr"] = GetRandomString(32)
		ret["package"] = "prepay_id=" + resp.PrepayId
		ret["signType"] = "MD5"
		sign := wxcore.Sign2(ret, serv.ApiKey, nil)
		ret["paySign"] = sign
	}
	return ret, nil
}

func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
