package payloaddelete

type SubscriptionRequest struct {
	SubID  string `json:"subscription_id" example:"123e4567-e89b-12d3-a456-426614174000"`
	UserID string `json:"user_id" example:"123e4567-e89b-12d3-a456-426614174000"`
}
