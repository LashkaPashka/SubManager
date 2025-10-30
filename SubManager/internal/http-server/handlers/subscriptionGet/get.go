package subscriptionget

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	payloadget "github.com/lashkapashka/SubManager/internal/http-server/handlers/subscriptionGet/payloadGet"
	"github.com/lashkapashka/SubManager/internal/model"
)

type Service interface {
	GetSubsByUserID(ctx context.Context, userID string) (subsModel []model.SubscriptionInputModel, err error)
}

// @Summary      Получить подписки
// @Description  Получает подписки конкретного пользователя из базы данных
// @Tags         subscriptions
// @Accept       json
// @Produce      json
// @Param        user_id  query  string  true "ID пользователя"
// @Success      200  {object}  payloadget.SubscriptionsResponse 
// @Failure      400  {string}  string  "Invalid request. Please check the submitted data."
// @Failure      500  {string}  string  "Internal server error"
// @Router       /subscriptions [get]
func Get(service Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("user_id")
		
		subsModel, err := service.GetSubsByUserID(r.Context(), userID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error!"))

			return 
		}

		resp := payloadget.SubscriptionsResponse{
			Subscriptions: subsModel,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&resp)
	}
}
