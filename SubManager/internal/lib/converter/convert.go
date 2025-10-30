package converter

import (
	"time"

	"github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionCreate/payloadcreate"
	"github.com/lashkapashka/SubManager/internal/model"
)

func Convert(paylod *payloadcreate.SubscriptionRequest) model.SubscriptionInputModel {
	startDate, _ := time.Parse("01-2006", paylod.StartDate)
	endDate, _ := time.Parse("01-2006", paylod.EndDate)

	return model.SubscriptionInputModel{
		ServiceName: paylod.ServiceName,
		Price: uint(paylod.Price),
		UserID: paylod.UserID,
		StartDate: startDate,
		EndDate: endDate,
	}
}