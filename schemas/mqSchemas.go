package schemas

type MqSchema struct {
	Header map[string]interface{} `json:"header"`
	Body   map[string]interface{} `json:"body"`
}

type MqStringSchema struct {
	Header string `json:"header"`
	Body   string `json:"body"`
}
