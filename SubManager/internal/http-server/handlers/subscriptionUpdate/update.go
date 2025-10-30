package subscriptionupdate

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	payloadupdate "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionUpdate/payloadUpdate"
	"github.com/lashkapashka/SubManager/internal/lib/req"
	"github.com/lashkapashka/SubManager/internal/model"
)

type Service interface {
	UpdateSubscription(ctx context.Context, subID, userID string) (subModel model.SubscriptionInputModel, err error)
}

// @Summary      Обновить подписку
// @Description  Обновляет цену подписки конкретного пользователя в базе данных
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription   body payloadupdate.SubscriptionRequest true "Данные для изменения"
// @Success      200  {object}  payloadupdate.SubscriptionResponse 
// @Failure      400  {string}  string  "Invalid request. Please check the submitted data."
// @Failure      500  {string}  string  "Internal server error"
// @Router       /subscriptions [patch]
func Update(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payloadupdate.SubscriptionRequest](w, r, logger)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Validation error"))

			return
		}

		ctx := context.WithValue(context.Background(), "price", body.Price)
		
		req := r.WithContext(ctx)

		subModel, err := service.UpdateSubscription(req.Context(), body.SubID, body.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error!"))

			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&payloadupdate.SubscriptionResponse{
			UserID: subModel.UserID,
			ServiceName: subModel.ServiceName,
			Price: int(subModel.Price),
		})
	}
}
