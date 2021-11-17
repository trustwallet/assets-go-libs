package processor

import (
	log "github.com/sirupsen/logrus"
	"github.com/trustwallet/assets-backend/internal/config"
	"github.com/trustwallet/assets-backend/internal/validators"
	"github.com/trustwallet/assets-backend/pkg/assetfs"
	"github.com/trustwallet/assets-backend/pkg/assetfs/validation"
)

type Service struct {
	conf              config.ValidatorsSettings
	fileStorage       *assetfs.FileProvider
	validatorsService *validators.Service
}

func NewService(
	conf config.ValidatorsSettings,
	fileStorage *assetfs.FileProvider,
	validatorsService *validators.Service,
) *Service {
	return &Service{conf: conf, fileStorage: fileStorage, validatorsService: validatorsService}
}

func (s *Service) RunSanityCheck(paths []string) error {
	//wg := &sync.WaitGroup{}

	for _, path := range paths {
		//wg.Add(1)

		//go func(path string) {
		//defer wg.Done()

		file, err := s.fileStorage.GetAssetFile(path)
		if err != nil {
			log.WithError(err).Error()
			return err
		}

		validator := s.validatorsService.GetValidatorForFile(file)

		if validator != nil {
			err = validator.Run(file)
			if err != nil {
				HandleError(err, file.Info, validator.ValidationName)
			}
		}

		err = file.Close()
		if err != nil {
			log.WithError(err).Error()

			return err
		}
	}
	//(path)
	//}
	//wg.Wait()

	return nil
}

func HandleError(err error, info *assetfs.AssetInfo, valName string) {
	errors := UnwrapComposite(err)

	for _, err := range errors {
		if warn, ok := err.(*validation.Warning); ok {
			//log.WithField("path", info.Path()).Warning(warn)

			HandleWarning(warn)

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
		//TODO errors handling call fixers
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

func HandleWarning(warning *validation.Warning) {

}
