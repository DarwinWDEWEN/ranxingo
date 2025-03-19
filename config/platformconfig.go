package config

type Platform struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	ChatURL string `json:"chat_url"`
	ImgURL  string `json:"img_url"`
}

var OpenAI = Platform{
	Name:    "OpenAI - GPT",
	Value:   "OpenAI",
	ChatURL: "https://api.chat-plus.net/v1/chat/completions",
	ImgURL:  "https://api.chat-plus.net/v1/images/generations",
}
var Azure = Platform{
	Name:    "微软 - Azure",
	Value:   "Azure",
	ChatURL: "https://chat-bot-api.openai.azure.com/openai/deployments/{model}/chat/completions?api-version=2023-05-15",
}
var ChatGLM = Platform{
	Name:    "智谱 - ChatGLM",
	Value:   "ChatGLM",
	ChatURL: "https://open.bigmodel.cn/api/paas/v3/model-api/{model}/sse-invoke",
}
var Baidu = Platform{
	Name:    "百度 - 文心大模型",
	Value:   "Baidu",
	ChatURL: "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/{model}",
}
var XunFei = Platform{
	Name:    "讯飞 - 星火大模型",
	Value:   "XunFei",
	ChatURL: "wss://spark-api.xf-yun.com/{version}/chat",
}
var QWen = Platform{
	Name:    "阿里 - 通义千问",
	Value:   "QWen",
	ChatURL: "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation",
}
