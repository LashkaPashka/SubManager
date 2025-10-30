package payloadget

import "github.com/lashkapashka/SubManager/internal/model"

type SubscriptionsResponse struct {
    Subscriptions []model.SubscriptionInputModel `json:"subscriptions"`
}