package validation

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func ValidateJson(b []byte) error {
	if !json.Valid(b) {
		return ErrInvalidJson
	}

	return nil
}

func validateDuplicateKeys(d *json.Decoder, path []string) error {
	mainToken, err := d.Token()
	if err != nil {
		return err
	}

	delimiter, ok := mainToken.(json.Delim)

	if !ok {
		return nil
	}

	if delimiter == '{' {
		keys := make(map[string]bool)
		for d.More() {
			theToken, err := d.Token()
			if err != nil {
				return err
			}

			key := theToken.(string)

			if _, exists := keys[key]; exists {
				return fmt.Errorf("duplicate key: %s", key)
			}
			keys[key] = true

			if err := validateDuplicateKeys(d, append(path, key)); err != nil {
				return fmt.Errorf("invalid value on key: %s", key)
			}

		}
		if _, err := d.Token(); err != nil {
			return err
		}
	} else if delimiter == '[' {
		counter := 0
		for d.More() {
			if err := validateDuplicateKeys(d, append(path, strconv.Itoa(counter))); err != nil {
				return err
			}

			counter++
		}

		if _, err := d.Token(); err != nil {
			return err
		}
	}

	return nil
}
