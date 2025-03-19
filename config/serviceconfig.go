package config

type ApiConfig struct {
	ApiURL string
	AppId  string
	Token  string
}

type MjProxyConfig struct {
	Enabled bool
	ApiURL  string // api 地址
	Mode    string // 绘画模式，可选值：fast/turbo/relax
	ApiKey  string
}

type StableDiffusionConfig struct {
	Enabled bool
	Model   string // 模型名称
	ApiURL  string
	ApiKey  string
}

type MjPlusConfig struct {
	Enabled bool   // 如果启用了 MidJourney Plus，将会自动禁用原生的MidJourney服务
	ApiURL  string // api 地址
	Mode    string // 绘画模式，可选值：fast/turbo/relax
	ApiKey  string
}
