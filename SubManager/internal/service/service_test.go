package service

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/lashkapashka/SubManager/internal/storage/postgresql"
)

var sv *Service
var connStr string = "postgres://postgres:root@localhost:5432/submanager"


func TestMain(m *testing.M) {
	logger := setupLogger("test")

	storage := postgresql.New(connStr, logger)
	
	sv = New(storage, logger)

	m.Run()
}

func TestTotalSubscription(t *testing.T) {
	date := "07-2025"
	mp := make(map[string]string, 1)
	mp["service_name"] = "NetFlix"

	totalSum, err := sv.TotalSubscription(context.Background(), date, mp)
	if err != nil {
		t.Fatal("Error 1")
	}

	if totalSum == 0 {
		t.Fatal("TotalSum equal 0")
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
