package subscriptioncreate

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionCreate/payloadcreate"
	"github.com/lashkapashka/SubManager/internal/lib/converter"
	"github.com/lashkapashka/SubManager/internal/lib/req"
	"github.com/lashkapashka/SubManager/internal/model"
)

type Service interface {
	CreateSubscription(ctx context.Context, subModel model.SubscriptionInputModel) (success string, err error)
}

// @Summary      Создать подписку
// @Description  Создаёт подписку в базе данных
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription body payloadcreate.SubscriptionRequest true "Название сервиса"
// @Success      200  {string}  string	"Subscription was created!"
// @Failure      400  {string}  string  "Invalid request. Please check the submitted data."
// @Failure      500  {string}  string  "Internal server error"
// @Router       /subscriptions [post]
func CreateSubscription(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Validation request-payload
		body, err := req.HandleBody[payloadcreate.SubscriptionRequest](w, r, logger)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Validation error"))
			
			return
		}

		// 2. Convert body request
		subModel := converter.Convert(&body)

		// 3. Service
		success, err := service.CreateSubscription(r.Context(), subModel)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error service"))
			
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(success))
	}
}