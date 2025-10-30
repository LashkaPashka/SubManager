package payloadprice

type ResponseDate struct {
    Month       string `json:"month" example:"09-2025"`
    TotalSum    int    `json:"total_sum" example:"400"`
    UserID      string `json:"user_id,omitempty" example:"60601fee-2bf1-4721-ae6f-7636e79a0cba"`
    ServiceName string `json:"service_name,omitempty" example:"Netflix"`
}