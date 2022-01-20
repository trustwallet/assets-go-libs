package list

import (
	"fmt"

	"github.com/trustwallet/assets-go-libs/validation"
)

func ValidateList(list []Model) error {
	compErr := validation.NewErrComposite()

	for _, validator := range list {
		if err := validateRequiredFields(validator); err != nil {
			compErr.Append(err)
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func validateRequiredFields(model Model) error {
	if model.Name == nil || model.ID == nil || model.Description == nil || model.Website == nil {
		return fmt.Errorf("%w: it must have more fields", validation.ErrMissingField)
	}

	return nil
}
