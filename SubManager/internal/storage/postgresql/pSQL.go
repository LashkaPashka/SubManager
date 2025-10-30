package postgresql

import (
	"context"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	workoption "github.com/lashkapashka/SubManager/internal/lib/workOption"
	"github.com/lashkapashka/SubManager/internal/model"
)

const (
	userID = "user_id"
	serviceName = "service_name"
)

type Storage struct {
	pool   *pgxpool.Pool
	logger *slog.Logger
}

func New(connStr string, logger *slog.Logger) *Storage {
	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		logger.Error("Invalid run Db")
		return nil
	}

	if err := pool.Ping(context.Background()); err != nil {
		logger.Error("Invalid ping Db")
		return nil
	}

	logger.Info("PostgreSQL's running", slog.Int("port", 5432))

	return &Storage{
		pool:   pool,
		logger: logger,
	}
}

func (s *Storage) Create(ctx context.Context, subModel model.SubscriptionInputModel) (subID string, err error) {
	const op = "SubManager.storage.Create"

	query := `INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
			  VALUES (@service_name, @price, @user_id, @start_date, @end_date)
			  RETURNING sub_id`

	args := pgx.NamedArgs{
		"service_name": subModel.ServiceName,
		"price":        subModel.Price,
		"user_id":      subModel.UserID,
		"start_date":   subModel.StartDate,
		"end_date":     subModel.EndDate,
	}

	if err = s.pool.QueryRow(ctx, query, args).Scan(&subID); err != nil {
		s.logger.Error("Invalid create payload",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)
		return subID, err
	}

	slog.Info("Record was created successfuly in Db")

	return subID, err
}

func (s *Storage) GetByUserID(ctx context.Context, userID string) (subsModel []model.SubscriptionInputModel, err error) {
	const op = "SubManager.storage.Get"

	query := `
		SELECT service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE user_id = $1
		ORDER BY price;
		`

	rows, err := s.pool.Query(ctx, query, userID)
	if err != nil {
		s.logger.Error("Invalid get payload",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)
		return subsModel, err
	}

	for rows.Next() {
		var subModel model.SubscriptionInputModel

		if err = rows.Scan(
			&subModel.ServiceName,
			&subModel.Price,
			&subModel.UserID,
			&subModel.StartDate,
			&subModel.EndDate,
		); err != nil {
			return subsModel, err
		}

		subsModel = append(subsModel, subModel)
	}

	s.logger.Info("Record was get successfully from Db")

	return subsModel, err
}

func (s *Storage) Update(ctx context.Context, subID, userID string) (serviceName string, price int, err error) {
	const op = "SubManager.storage.Update"

	query := `UPDATE subscriptions
			  SET price = @price
			  WHERE sub_id = @sub_id AND user_id = @user_id
			  RETURNING service_name, price
			 `

	args := pgx.NamedArgs{
		"price": ctx.Value("price"),
		"sub_id": subID,
		"user_id": userID,
	}

	if err = s.pool.QueryRow(ctx, query, args).Scan(&serviceName, &price); err != nil {
		s.logger.Error("Invalid get payload",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return serviceName, price, err
	}

	s.logger.Info("Payload was updated successfully in Db")

	return serviceName, price, err
}

func (s *Storage) Delete(ctx context.Context, subID, userID string) (string, error) {
	const op = "SubManager.storage.Delete"

	query := `DELETE FROM subscriptions
			  WHERE sub_id = @sub_id AND user_id = @user_id
			 `

	args := pgx.NamedArgs{
		"sub_id":  subID,
		"user_id": userID,
	}

	if _, err := s.pool.Exec(ctx, query, args); err != nil {
		s.logger.Error("Invalid get payload",
			slog.String("op", op),
			slog.String("err", err.Error()),
		)

		return "", err
	}

	s.logger.Info("Payload was deleted successfully from Db")

	return subID, nil
}

func (s *Storage) Total(ctx context.Context, date, key, value string) (totalSum int, err error) {
	const op = "SubManager.storage.Total"
	
	var (
		query string
		args  pgx.NamedArgs
	)

	switch key {
		case userID:
			s.logger.Info("Calculating total by user ID", slog.String("user_id", value))
			query, args = workoption.WorkUserID(date, value)

		case serviceName:
			s.logger.Info("Calculating total by service name", slog.String("service_name", value))
			query, args = workoption.WorkServiceName(date, value)

		default:
			s.logger.Info("Calculating total by date only", slog.String("date", date))
			query, args = workoption.WorkDate(date)
	}

	err = s.pool.QueryRow(ctx, query, args).Scan(&totalSum)
	if err != nil {
		s.logger.Error("failed to calculate total",
			slog.String("op", op),
			slog.String("query", query),
			slog.String("err", err.Error()),
		)
		return 0, err
	}

	return totalSum, nil
}