package config

type SystemConfig struct {
	Title         string `json:"title,omitempty"`       // 网站标题
	Slogan        string `json:"slogan,omitempty"`      // 网站 slogan
	AdminTitle    string `json:"admin_title,omitempty"` // 管理后台标题
	Logo          string `json:"logo,omitempty"`
	InitPower     int    `json:"init_power,omitempty"`      // 新用户注册赠送算力值
	DailyPower    int    `json:"daily_power,omitempty"`     // 每日赠送算力
	InvitePower   int    `json:"invite_power,omitempty"`    // 邀请新用户赠送算力值
	VipMonthPower int    `json:"vip_month_power,omitempty"` // VIP 会员每月赠送的算力值

	RegisterWays    []string `json:"register_ways,omitempty"`    // 注册方式：支持手机（mobile），邮箱注册（email），账号密码注册
	EnabledRegister bool     `json:"enabled_register,omitempty"` // 是否开放注册

	RewardImg     string  `json:"reward_img,omitempty"`     // 众筹收款二维码地址
	EnabledReward bool    `json:"enabled_reward,omitempty"` // 启用众筹功能
	PowerPrice    float64 `json:"power_price,omitempty"`    // 算力单价

	OrderPayTimeout int    `json:"order_pay_timeout,omitempty"` //订单支付超时时间
	VipInfoText     string `json:"vip_info_text,omitempty"`     // 会员页面充值说明
	DefaultModels   []int  `json:"default_models,omitempty"`    // 默认开通的 AI 模型

	MjPower       int `json:"mj_power,omitempty"`        // MJ 绘画消耗算力
	MjActionPower int `json:"mj_action_power,omitempty"` // MJ 操作（放大，变换）消耗算力
	SdPower       int `json:"sd_power,omitempty"`        // SD 绘画消耗算力
	DallPower     int `json:"dall_power,omitempty"`      // DALLE3 绘图消耗算力

	WechatCardURL string `json:"wechat_card_url,omitempty"` // 微信客服地址

	EnableContext bool `json:"enable_context,omitempty"`
	ContextDeep   int  `json:"context_deep,omitempty"`

	SdNegPrompt string `json:"sd_neg_prompt"` // SD 默认反向提示词

	IndexBgURL string `json:"index_bg_url"` // 前端首页背景图片
}
