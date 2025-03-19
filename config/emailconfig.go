package config

type SmtpConfig struct {
	UseTls   bool // 是否使用 TLS 发送
	Host     string
	Port     int
	AppName  string // 应用名称
	From     string // 发件人邮箱地址
	Password string // 发件人邮箱密码
}
