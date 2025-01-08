package models

type Payload struct {
	Username    string           `json:"username"`
	Text        string           `json:"text"`
	IconURL     string           `json:"icon_url"`
	Attachments []map[string]any `json:"attachments"`
}
