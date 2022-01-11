package validation

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

func ValidateJson(b []byte) error {
	if !json.Valid(b) {
		return ErrInvalidJson
	}

	return checkDuplicateKey(json.NewDecoder(bytes.NewReader(b)), nil)
}

func checkDuplicateKey(d *json.Decoder, path []string) error {
	// Get next token from JSON.
	t, err := d.Token()
	if err != nil {
		return err
	}

	delim, ok := t.(json.Delim)
	if !ok {
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]bool)
		for d.More() {
			// Get attribute key.
			t, err := d.Token()
			if err != nil {
				return err
			}
			key := t.(string)

			// Check for duplicates.
			if keys[key] {
				return fmt.Errorf("duplicate key '%s'", strings.Join(append(path, key), "/"))
			}
			keys[key] = true

			// Check value.
			if err := checkDuplicateKey(d, append(path, key)); err != nil {
				return err
			}
		}

		// Consume trailing "}".
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := checkDuplicateKey(d, append(path, strconv.Itoa(i))); err != nil {
				return err
			}

			i++
		}

		// Consume trailing "]".
		if _, err := d.Token(); err != nil {
			return err
		}
	}

	return nil
}
