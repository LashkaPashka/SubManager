package payloaddelete

type SubscriptionRequest struct {
	SubID  string `json:"subscription_id" example:"3234234"`
	UserID string `json:"user_id" example:"u23423423uid"`
}
