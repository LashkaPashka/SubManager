package subscriptiontotal

import (
	"context"
	"log/slog"
	"net/http"

	payloadprice "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionTotalPrice/payloadPrice"
	"github.com/lashkapashka/SubManager/internal/lib/res"
)

type Service interface {
	TotalSubscription(ctx context.Context, date string, mp map[string]string) (totalSum int, err error)
}

// @Summary      Получить итоговую сумму
// @Description  Возвращает итоговую сумму по переданным параметрам: только по дате, по дате и userID или по дате и названию сервиса.
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        month        query  string true "Месяц и год (MM-YYYY)"
// @Param        user_id      query  string false "ID пользователя (UUID)"
// @Param        service_name query  string false "Название сервиса (например, Netflix)"
// @Success      200  {object}  payloadprice.ResponseDate
// @Failure      400  {string}  string  "Invalid request. Please check the submitted data."
// @Failure      500  {string}  string  "Internal server error"
// @Router       /subscriptions/total [get]
func TotalPrice(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var mp map[string]string = make(map[string]string, 2)

		query := r.URL.Query()
		date := query.Get("month")

		switch {
			case query.Has("user_id"):
				mp["user_id"] = query.Get("user_id")

			case query.Has("service_name"):
				mp["service_name"] = query.Get("service_name")
		}
		
		if date == "" {
			http.Error(w, "missing required parameter: date", http.StatusBadRequest)
			return
		}

		totalSum, err := service.TotalSubscription(r.Context(), date, mp)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error!"))

			return
		}

		dateResp := payloadprice.ResponseDate{
			Month: date,
			TotalSum: totalSum,
		}

		if userID, ok := mp["user_id"]; ok {
			dateResp.UserID = userID
		}

		if serviceName, ok := mp["service_name"]; ok {
			dateResp.ServiceName = serviceName
		}

		res.Encode(w, &dateResp)
	}
}