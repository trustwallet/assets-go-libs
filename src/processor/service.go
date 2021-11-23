package processor

import (
	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/src/reporter"
	"github.com/trustwallet/assets-go-libs/src/validator"
)

const (
	reportSanityCheckKey = "sanity-check"
)

type Service struct {
	fileService       *file.FileService
	validatorsService *validator.Service
	reporterService   *reporter.Service
}

func NewService(fs *file.FileService, vs *validator.Service, rs *reporter.Service) *Service {
	return &Service{
		fileService:       fs,
		validatorsService: vs,
		reporterService:   rs,
	}
}

func (s *Service) RunSanityCheck(paths []string) error {
	report := s.reporterService.GetOrNew(reportSanityCheckKey)

	for _, path := range paths {
		f, err := s.fileService.GetAssetFile(path)
		if err != nil {
			log.WithError(err).Error()
			return err
		}

		report.TotalFiles += 1

		validator := s.validatorsService.GetFoldersFilesValidator(f)
		if validator != nil {
			err = validator.Run(f)
			if err != nil {
				HandleError(err, f.Info, validator.ValidationName, report)
			}
		}

		err = f.Close()
		if err != nil {
			log.WithError(err).Error()
			return err
		}
	}

	err := s.reporterService.Update(reportSanityCheckKey, report)
	if err != nil {
		log.WithError(err).Error()
		return err
	}

	return nil
}

func HandleError(err error, info *file.AssetInfo, valName string, report *reporter.Report) {
	errors := UnwrapComposite(err)

	for _, err := range errors {
		if warn, ok := err.(*validation.Warning); ok {
			report.Warnings += 1
			HandleWarning(warn, info)
			continue
		} else {
			report.Errors += 1

			log.WithField("type", info.Type()).
				WithField("chain", info.Chain().Handle).
				WithField("asset", info.Asset()).
				WithField("path", info.Path()).
				WithField("validation", valName).
				Error(err)
		}

		switch err {
		// TODO: Call fixers here.
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
