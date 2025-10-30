package req

import (
	"log/slog"
	"net/http"
)


func HandleBody[T any](w http.ResponseWriter, r *http.Request, logger *slog.Logger) (T, error) {
	body, err := decode[T](r.Body, logger)
	if err != nil {
		return body, err
	}

	if err := isValid(body, logger); err != nil {
		return body, err
	}

	return body, nil
}