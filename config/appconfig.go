package config

// TODO 转移到app端
type AppConfig struct {
	Path           string `toml:"-"`
	Listen         string
	Session        Session
	AdminSession   Session
	ProxyURL       string
	MysqlDns       Mysqldns                // mysql 连接地址
	StaticDir      string                  // 静态资源目录
	StaticUrl      string                  // 静态资源 URL
	Redis          RedisConfig             // redis 连接信息
	ApiConfig      ApiConfig               // ChatPlus API authorization configs
	SMS            SMSConfig               // send mobile message config
	OSS            OSSConfig               // OSS config
	MjProxyConfigs []MjProxyConfig         // MJ proxy config
	MjPlusConfigs  []MjPlusConfig          // MJ plus config
	WeChatBot      bool                    // 是否启用微信机器人
	SdConfigs      []StableDiffusionConfig // sd AI draw service pool

	AlipayConfig    AlipayConfig    // 支付宝支付渠道配置
	HuPiPayConfig   HuPiPayConfig   // 虎皮椒支付配置
	SmtpConfig      SmtpConfig      // 邮件发送配置
	JPayConfig      JPayConfig      // payjs 支付配置
	WechatPayConfig WechatPayConfig // 微信支付渠道配置
	TikaHost        string          // TiKa 服务器地址
}
