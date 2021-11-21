package list

import (
	"fmt"

	"github.com/trustwallet/assets-go-libs/pkg/validation"
)

func ValidateList(list []Model) error {
	compErr := validation.NewErrComposite()
	for _, l := range list {
		if err := validateRequiredFields(&l); err != nil {
			compErr.Append(err)
		}
	}

	if compErr.Len() > 0 {
		return compErr
	}

	return nil
}

func validateRequiredFields(model *Model) error {
	if model.Name == nil || model.ID == nil || model.Description == nil || model.Website == nil {
		return fmt.Errorf("missing required fields")
	}

	return nil
}
