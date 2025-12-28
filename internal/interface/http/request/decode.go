package request

import (
	"belscourrsego/internal/interface/http/httperr"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func DecodeJSON[T any](r *http.Request, dst *T) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(dst); err != nil {
		if errors.Is(err, io.EOF) {
			return httperr.NewCodeInvalidInput(nil, err)
		}

		return httperr.NewCodeInvalidInput(nil, err)
	}

	return nil
}
