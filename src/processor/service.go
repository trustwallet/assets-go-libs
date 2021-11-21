package processor

import (
	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/src/validator"
)

type Service struct {
	fileStorage       *file.FileProvider
	validatorsService *validator.Service
}

func NewService(storage *file.FileProvider, service *validator.Service) *Service {
	return &Service{
		fileStorage:       storage,
		validatorsService: service,
	}
}

func (s *Service) RunSanityCheck(paths []string) error {
	for _, path := range paths {
		f, err := s.fileStorage.GetAssetFile(path)
		if err != nil {
			log.WithError(err).Error()
			return err
		}

		validator := s.validatorsService.GetValidatorForFile(f)
		if validator != nil {
			err = validator.Run(f)
			if err != nil {
				HandleError(err, f.Info, validator.ValidationName)
			}
		}

		err = f.Close()
		if err != nil {
			log.WithError(err).Error()

			return err
		}
	}

	return nil
}

func HandleError(err error, info *file.AssetInfo, valName string) {
	errors := UnwrapComposite(err)

	for _, err := range errors {
		if warn, ok := err.(*validation.Warning); ok {
			HandleWarning(warn, info)
			continue
		} else {
			log.WithField("file_type", info.Type()).
				WithField("file_chain", info.Chain()).
				WithField("file_asset", info.Asset()).
				WithField("file_path", info.Path()).
				WithField("validation_name", valName).
				Errorf("%+v", err)
		}

		switch err {
		// TODO: errors handling call fixers
		}
	}
}

func UnwrapComposite(err error) []error {
	compErr, ok := err.(*validation.ErrComposite)
	if !ok {
		return []error{err}
	}

	var errors []error
	for _, e := range compErr.GetErrors() {
		errors = append(errors, UnwrapComposite(e)...)
	}

	return errors
}

func HandleWarning(warning *validation.Warning, info *file.AssetInfo) {
	log.WithField("path", info.Path()).Warning(warning)
}
