package service

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/lashkapashka/SubManager/internal/model"
)

type Storage interface {
	Create(ctx context.Context, subModel model.SubscriptionInputModel) (subID string, err error)
	GetByUserID(ctx context.Context, userID string) (subsModel []model.SubscriptionInputModel, err error)
	Update(ctx context.Context, subID, userID string) (serviceName string, price int, err error)
	Delete(ctx context.Context, subID, userID string) (string, error)
	Total(ctx context.Context, date, key, value string) (totalSum int, err error)
}

type Service struct {
	logger  *slog.Logger
	storage Storage
}

func New(storage Storage, logger *slog.Logger) *Service {
	return &Service{
		logger:  logger,
		storage: storage,
	}
}

func (s *Service) CreateSubscription(ctx context.Context, subModel model.SubscriptionInputModel) (success string, err error) {
	const op = "SubManager.service.CreateSubscription"

	subID, err := s.storage.Create(ctx, subModel)
	if err != nil {
		s.logger.Error("Invalid service createSubscription",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return "Failed", err
	}

	success = fmt.Sprintf("Record:%s - created successfully", subID)

	return success, err
}

func (s *Service) GetSubsByUserID(ctx context.Context, userID string) (subsModel []model.SubscriptionInputModel, err error) {
	const op = "SubManager.service.GetSubByUserID"

	subsModel, err = s.storage.GetByUserID(ctx, userID)
	if err != nil {
		s.logger.Error("Invalid service GetSubByUserID",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return subsModel, err
	}

	return subsModel, err
}

func (s *Service) UpdateSubscription(ctx context.Context, subID, userID string) (subModel model.SubscriptionInputModel, err error) {
	const op = "SubManager.service.UpdateSubscription"

	serviceName, price, err := s.storage.Update(ctx, subID, userID)
	if err != nil {
		s.logger.Error("Invalid service UpdateSubscription",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return subModel, err
	}

	subModel = model.SubscriptionInputModel{
		ServiceName: serviceName,
		Price:       uint(price),
		UserID:      userID,
	}

	return subModel, err
}

func (s *Service) DeleteSubscription(ctx context.Context, subID, userID string) (success string, err error) {
	const op = "SubManager.service.DeleteSubscription"

	subID, err = s.storage.Delete(ctx, subID, userID)
	if err != nil {
		s.logger.Error("Invalid service DeleteSubscription",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return success, err
	}

	success = fmt.Sprintf("Record:%s - deleted successfully", subID)

	return success, err
}

func (s *Service) TotalSubscription(ctx context.Context, date string, mp map[string]string) (totalSum int, err error)  {
	const op = "SubManager.service.TotalSubscription"
	var key, value string

	if mp == nil {
		key = "date"
	} else {
		for _, k := range []string{"user_id", "service_name"} {
			if v, ok := mp[k]; ok {
				key, value = k, v
				break
			}
		}
	}

	fmt.Printf("DEBUG → key: %s | value: %s | date: %s\n", key, value, date)

	totalSum, err = s.storage.Total(ctx, date, key, value)
	if err != nil {
		s.logger.Error("Invalid service TotalSubscription",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)
		return 0, err
	}

	return totalSum, nil
}