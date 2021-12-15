package processor

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/src/core"
)

type Service struct {
	fileService *file.Service
	coreService *core.Service
}

func NewService(fs *file.Service, cs *core.Service) *Service {
	return &Service{
		fileService: fs,
		coreService: cs,
	}
}

func (s *Service) RunJob(paths []string, job func(*file.AssetFile)) error {
	for _, path := range paths {
		f, err := s.fileService.GetAssetFile(path)
		if err != nil {
			return fmt.Errorf("failed to get asset file: %s", err)
		}

		job(f)

		if err = f.Close(); err != nil {
			return fmt.Errorf("failed to close asset file: %s", err)
		}
	}

	return nil
}

func (s *Service) Check(f *file.AssetFile) {
	validator := s.coreService.GetValidator(f)

	if validator != nil {
		if err := validator.Run(f); err != nil {
			// TODO: somehow return an error from Check if there are any errors.
			HandleError(err, f.Info, validator.Name)
		}
	}
}

func (s *Service) Fix(f *file.AssetFile) {
	fixers := s.coreService.GetFixers(f)

	for _, fixer := range fixers {
		if err := fixer.Run(f); err != nil {
			HandleError(err, f.Info, fixer.Name)
		}
	}
}

func (s *Service) RunUpdateAuto() error {
	updaters := s.coreService.GetUpdatersAuto()

	for _, updater := range updaters {
		err := updater.Run()
		if err != nil {
			log.WithError(err).Error()
		}
	}

	return nil
}

func HandleError(err error, info *file.AssetInfo, valName string) {
	errors := UnwrapComposite(err)

	for _, err := range errors {
		logFields := log.Fields{
			"type":       info.Type(),
			"chain":      info.Chain().Handle,
			"asset":      info.Asset(),
			"path":       info.Path(),
			"validation": valName,
		}

		if warn, ok := err.(*validation.Warning); ok {
			log.WithFields(logFields).Warning(warn)
		} else {
			log.WithFields(logFields).Error(err)
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
