package validation

import (
	"encoding/json"
)

func ValidateJson(b []byte) error {
	if !json.Valid(b) {
		return ErrInvalidJson
	}

	return nil
}
