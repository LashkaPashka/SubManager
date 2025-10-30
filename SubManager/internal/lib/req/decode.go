package req

import (
	"encoding/json"
	"io"
	"log/slog"
)

func decode[T any](body io.ReadCloser, logger *slog.Logger) (T, error) {
	var payload T

	err := json.NewDecoder(body).Decode(&payload)

	if err != nil {
		logger.Error("Invalid decode body request", 
			slog.String("err", err.Error()),
		)
		return payload, err
	}

	return payload, err
}