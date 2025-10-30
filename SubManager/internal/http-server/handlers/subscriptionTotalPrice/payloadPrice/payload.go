package payloadprice

type ResponseDate struct {
    Month       string `json:"month" example:"09-2025"`
    TotalSum    int    `json:"total_sum" example:"400"`
    UserID      string `json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174000"`
    ServiceName string `json:"service_name,omitempty" example:"Netflix"`
}