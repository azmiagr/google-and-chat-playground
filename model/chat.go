package model

type SendMessageInput struct {
	UserID  int    `json:"user_id"`
	Content string `json:"content"`
}

type CreateConversationInput struct {
	UserIDs []int  `json:"user_ids"`
	Title   string `json:"title"`
}
