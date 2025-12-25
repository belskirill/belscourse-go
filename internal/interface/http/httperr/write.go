package httperr

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

func StatusFromCode(code Code) int {
	switch code {
	case CodeInternal:
		return http.StatusInternalServerError
	case CodeConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func Write(w http.ResponseWriter, logger *zap.Logger, httpError *HTTPError) {
	statusCode := StatusFromCode(httpError.Code)

	if statusCode == http.StatusInternalServerError {
		logger.Error("request failed",
			zap.Error(httpError.Err),
		)
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(map[string]any{
		"error": httpError,
	})
}
