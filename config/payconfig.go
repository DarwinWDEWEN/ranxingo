package config

type AlipayConfig struct {
	Enabled         bool   // 是否启用该支付通道
	SandBox         bool   // 是否沙盒环境
	AppId           string // 应用 ID
	UserId          string // 支付宝用户 ID
	PrivateKey      string // 用户私钥文件路径
	PublicKey       string // 用户公钥文件路径
	AlipayPublicKey string // 支付宝公钥文件路径
	RootCert        string // Root 秘钥路径
	NotifyURL       string // 异步通知回调
	ReturnURL       string // 支付成功返回地址
}

type WechatPayConfig struct {
	Enabled    bool   // 是否启用该支付通道
	AppId      string // 公众号的APPID,如：wxd678efh567hg6787
	MchId      string // 直连商户的商户号，由微信支付生成并下发
	SerialNo   string // 商户证书的证书序列号
	PrivateKey string // 用户私钥文件路径
	ApiV3Key   string // API V3 秘钥
	NotifyURL  string // 异步通知回调
	ReturnURL  string // 支付成功返回地址
}

type HuPiPayConfig struct { //虎皮椒第四方支付配置
	Enabled   bool   // 是否启用该支付通道
	Name      string // 支付名称，如：wechat/alipay
	AppId     string // App ID
	AppSecret string // app 密钥
	ApiURL    string // 支付网关
	NotifyURL string // 异步通知回调
	ReturnURL string // 支付成功返回地址
}

// JPayConfig PayJs 支付配置
type JPayConfig struct {
	Enabled    bool
	Name       string // 支付名称，默认 wechat
	AppId      string // 商户 ID
	PrivateKey string // 私钥
	ApiURL     string // API 网关
	NotifyURL  string // 异步回调地址
	ReturnURL  string // 支付成功返回地址
}
