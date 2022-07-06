package client

type UpdateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Client `json:"result"`
}

type Update struct {
	ID      int    `json:"message_id"`
	Message string `json:"message"`
}