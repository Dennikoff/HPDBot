package client

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"message_id"`
	Message string `json:"message"`
}

type Message struct {
	ID   int    `json:"chat_id"`
	Text string `json:"text"`
}
