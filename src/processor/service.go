package processor

import (
	log "github.com/sirupsen/logrus"

	"github.com/trustwallet/assets-go-libs/pkg/file"
	"github.com/trustwallet/assets-go-libs/pkg/validation"
	"github.com/trustwallet/assets-go-libs/src/core"
	"github.com/trustwallet/assets-go-libs/src/reporter"
)

const reportSanityCheckKey = "sanity-check"

type Service struct {
	fileService     *file.Service
	coreService     *core.Service
	reporterService *reporter.Service
}

func NewService(fs *file.Service, cs *core.Service, rs *reporter.Service) *Service {
	return &Service{
		fileService:     fs,
		coreService:     cs,
		reporterService: rs,
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

		validator := s.coreService.GetValidator(f)
		fixer := s.coreService.GetFixer(f)

		if validator != nil {
			err = validator.Run(f)
			if err != nil {
				HandleError(err, f.Info, validator.Name, report)
			}
		}

		if fixer != nil && err != nil {
			err = fixer.Run(f)
			if err != nil {
				log.WithError(err).Error()
			} else {
				report.Fixed += 1
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

func (s *Service) RunUpdateAuto() error {
	updaters := s.coreService.GetUpdatersAuto()

	for _, updater := range updaters {
		log.WithField("Name", updater.Name).Debug("Running updater")

		err := updater.Run()
		if err != nil {
			log.WithError(err).Error()
		}
	}

	return nil
}

func HandleError(err error, info *file.AssetInfo, valName string, report *reporter.Report) {
	errors := UnwrapComposite(err)

	for _, err := range errors {
		if warn, ok := err.(*validation.Warning); ok {
			report.Warnings += 1
			log.WithField("path", info.Path()).Warning(warn)
		} else {
			report.Errors += 1
			log.WithFields(log.Fields{
				"type":       info.Type(),
				"chain":      info.Chain().Handle,
				"asset":      info.Asset(),
				"path":       info.Path(),
				"validation": valName,
			}).Error(err)
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
