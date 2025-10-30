package payloadupdate

type SubscriptionRequest struct {
	SubID  string `json:"sub_id"`
	UserID string `json:"user_id"`
	Price  int    `json:"price"`
}

type SubscriptionResponse struct {
	UserID      string `json:"user_id"`
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
}
