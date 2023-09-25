package models

type LoginResponse struct {
	Openid      string `json:"openid"`
	SessionKey  string `json:"session_key"`
	AccessToken string `json:"access_token"`
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
}

// TemplateMessage 模板消息
type TemplateMessage struct {
	AccessToken string `json:"access_token"`
	Appid       string `json:"appid"`
	Touser      string `json:"touser"`
	TemplateId  string `json:"templateId"`
	FormId      string `json:"formId"`
}
