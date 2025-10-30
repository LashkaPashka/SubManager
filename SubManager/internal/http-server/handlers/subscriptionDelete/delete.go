package subscriptiondelete

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	payloaddelete "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionDelete/payloadDelete"
	"github.com/lashkapashka/SubManager/internal/lib/req"
)

type Service interface {
	DeleteSubscription(ctx context.Context, subID, userID string) (success string, err error)
}

// @Summary      Удалить подписку
// @Description  Удаляет подписку в базе данных
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        subscription body payloaddelete.SubscriptionRequest true "Данные сервиса"
// @Success      200  {string}  string	"Subscription was deleted!"
// @Failure      400  {string}  string  "Invalid request. Please check the submitted data."
// @Failure      500  {string}  string  "Internal server error"
// @Router       /subscriptions [delete]
func Delete(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[payloaddelete.SubscriptionRequest](w, r, logger)
		if err != nil {
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte("Validation error"))

			return
		}

		success, err := service.DeleteSubscription(r.Context(), body.SubID, body.UserID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error!"))

			return 
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Type-Content", "application/json")
		json.NewEncoder(w).Encode(&success)
	}
}