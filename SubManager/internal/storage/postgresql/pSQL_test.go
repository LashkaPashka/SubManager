package postgresql

import (
	"context"
	"log/slog"
	"os"
	"testing"
)

var st *Storage

var connStr string = "postgres://postgres:root@localhost:5432/submanager"

func TestMain(m *testing.M) {
	logger := setupLogger("test")

	st = New(connStr, logger)

	m.Run()
}

func TestGetByUserID(t *testing.T) {
	userID := "60601fee-2bf1-4721-ae6f-7636e79a0cba"

	subsModel, err := st.GetByUserID(context.Background(), userID)
	if err != nil {
		t.Fatalf("Error! getByUserID, err - %v", err)
	}

	if len(subsModel) == 0 || subsModel == nil {
		t.Fatal("subsModel is empty")
	}

	t.Log(subsModel)
}

func TestUpdate(t *testing.T) {
	sub_id := "3a44a3fa-7762-4d57-b502-840b41b9c722"
	user_id := "60601fee-2bf1-4721-ae6f-7636e79a0cba"
	price := 2000

	ctx := context.WithValue(context.Background(), "price", price)

	if ctx.Value("price") == "" {
		t.Fatal("Key price is empty")
	}

	serviceName, price, err := st.Update(ctx, sub_id, user_id)
	if err != nil {
		t.Fatalf("Error: %v", err)
	}

	t.Log(serviceName, price)
}

func TestTotal(t *testing.T) {
	date := "07-2025"
	key := "service_name"
	value := "NetFlix"
	
	totalSum, err := st.Total(context.Background(), date, key, value)
	if err != nil {
		t.Fatal("Error 1")
	}

	if totalSum == 0 {
		t.Fatal("Sum equal 0")
	}

	t.Log(totalSum)
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "test":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}