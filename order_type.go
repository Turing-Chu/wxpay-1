package wxpay

import (
	"encoding/json"
)

const (
	K_TRADE_TYPE_JSAPI  = "JSAPI"
	K_TRADE_TYPE_NATIVE = "NATIVE"
	K_TRADE_TYPE_APP    = "APP"
)

// https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type UnifiedOrderParam struct {
	NotifyURL      string     `xml:"notify_url"`       // 是 异步接收微信支付结果通知的回调地址，通知url必须为外网可访问的url，不能携带参数。
	Body           string     `xml:"body"`             // 是 商品简单描述，该字段请按照规范传递，具体请见参数规定
	OutTradeNo     string     `xml:"out_trade_no"`     // 是 商户系统内部订单号，要求32个字符内，只能是数字、大小写字母_-|*@ ，且在同一个商户号下唯一。详见商户订单号
	TotalFee       int        `xml:"total_fee"`        // 是 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string     `xml:"spbill_create_ip"` // 是 APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	TradeType      string     `xml:"trade_type"`       // 是 取值如下：JSAPI，NATIVE，APP等，说明详见参数规定
	SignType       string     `xml:"sign_type"`        // 否 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	DeviceInfo     string     `xml:"device_info"`      // 否 自定义参数，可以为终端设备号(门店号或收银设备ID)，PC网页或公众号内支付可以传"WEB"
	Detail         string     `xml:"detail"`           // 否 商品详细描述，对于使用单品优惠的商户，改字段必须按照规范上传，详见“单品优惠参数说明”
	Attach         string     `xml:"attach"`           // 否 附加数据，在查询API和支付通知中原样返回，可作为自定义参数使用。
	FeeType        string     `xml:"fee_type"`         // 否 符合ISO 4217标准的三位字母代码，默认人民币：CNY，详细列表请参见货币类型
	TimeStart      string     `xml:"time_start"`       // 否 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire     string     `xml:"time_expire"`      // 否 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则  注意：最短失效时间间隔必须大于5分钟
	GoodsTag       string     `xml:"goods_tag"`        // 否 订单优惠标记，使用代金券或立减优惠功能时需要的参数，说明详见代金券或立减优惠
	ProductId      string     `xml:"product_id"`       // 否 trade_type=NATIVE时（即扫码支付），此参数必传。此参数为二维码中包含的商品ID，商户自行定义。
	LimitPay       string     `xml:"limit_pay"`        // 否 上传此参数no_credit--可限制用户不能使用信用卡支付
	OpenId         string     `xml:"open_id"`          // 否 trade_type=JSAPI时（即公众号支付），此参数必传，此参数为微信用户在商户对应appid下的唯一标识。openid如何获取，可参考【获取openid】。企业号请使用【企业号OAuth2.0接口】获取企业号内成员userid，再调用【企业号userid转openid接口】进行转换
	SceneInfo      string     `xml:"scene_info"`       // 否 该字段用于上报场景信息，目前支持上报实际门店信息。该字段为JSON对象数据，对象格式为{"store_info":{"id": "门店ID","name": "名称","area_code": "编码","address": "地址" }} ，字段详细说明请点击行前的+展开
	StoreInfo      *StoreInfo `xml:"-"`
}

type StoreInfo struct {
	Id       string `json:"id"`        // 门店唯一标识
	Name     string `json:"name"`      // 门店名称
	AreaCode string `json:"area_code"` // 门店所在地行政区划码，详细见《最新县及县以上行政区划代码》
	Address  string `json:"address"`   // 门店详细地址
}

func (this *UnifiedOrderParam) Params() map[string]interface{} {
	var m = make(map[string]interface{})
	m["notify_url"] = this.NotifyURL
	if len(this.SignType) == 0 {
		this.SignType = K_SIGN_TYPE_MD5
	}
	m["sign_type"] = this.SignType
	m["device_info"] = this.DeviceInfo
	m["body"] = this.Body
	m["detail"] = this.Detail
	m["attach"] = this.Attach
	m["out_trade_no"] = this.OutTradeNo
	m["fee_type"] = this.FeeType
	m["total_fee"] = this.TotalFee
	m["spbill_create_ip"] = this.SpbillCreateIP
	m["time_start"] = this.TimeStart
	m["time_expire"] = this.TimeExpire
	m["goods_tag"] = this.GoodsTag
	if len(this.TradeType) == 0 {
		this.TradeType = K_TRADE_TYPE_APP
	}
	m["trade_type"] = this.TradeType
	m["product_id"] = this.ProductId
	m["limit_pay"] = this.LimitPay
	m["open_id"] = this.OpenId

	if this.StoreInfo != nil {
		var storeInfoByte, err = json.Marshal(this.StoreInfo)
		if err == nil {
			this.SceneInfo = "{\"store_info\" :" + string(storeInfoByte) + "}"
			m["scene_info"] = this.SceneInfo
		}
	}
	return m
}

type UnifiedOrderResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	AppId      string `xml:"appid"`
	MCHId      string `xml:"mch_id"`
	DeviceInfo string `xml:"device_info"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	ResultCode string `xml:"result_code"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
	CodeURL    string `xml:"code_url"`
}
