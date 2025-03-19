package captcha

import (
	"github.com/mojocn/base64Captcha"
)

type CaptchaType string

const (
	CaptchaTypeAudio   CaptchaType = "Audio"
	CaptchaTypeString  CaptchaType = "String"
	CaptchaTypeChinese CaptchaType = "Chinese"
	CaptchaTypeMath    CaptchaType = "Math"
	CaptchaTypeDigit   CaptchaType = "Digit"
)

type CaptchaRes struct {
	Id      string
	B64Data string
	Asr     string
}

func GenerateCaptcha(capReq CaptchaType) (res CaptchaRes, err error) {
	var driver base64Captcha.Driver
	switch capReq {
	case CaptchaTypeAudio:
		driver = base64Captcha.DefaultDriverAudio
	case CaptchaTypeDigit:
		driver = base64Captcha.NewDriverDigit(34, 120, 6, 0.7, 80)
	case CaptchaTypeMath:
		driver = base64Captcha.DefaultDriverDigit
	case CaptchaTypeChinese:
		driver = base64Captcha.DefaultDriverDigit
	case CaptchaTypeString:
		driver = base64Captcha.DefaultDriverDigit
	default:
		driver = base64Captcha.DefaultDriverDigit
	}

	id, content, asr := driver.GenerateIdQuestionAnswer()
	item, err := driver.DrawCaptcha(content)
	b64s := item.EncodeB64string()
	if err != nil {
		return CaptchaRes{}, err
	}
	return CaptchaRes{
		Id:      id,
		B64Data: b64s,
		Asr:     asr,
	}, nil
}
