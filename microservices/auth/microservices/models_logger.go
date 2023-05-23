package microservices

type LoggerSlackRespTemplateDTO struct {
	Status int    `json:"status"`
	Reason string `json:"reason"`
}

type LoggerSlackTemplateDTO struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
