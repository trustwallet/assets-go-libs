package reporter

import (
	"fmt"
	"sync"
)

type Service struct {
	mu      *sync.Mutex
	reports ReportList
}

func NewReportService() *Service {
	return &Service{
		mu:      &sync.Mutex{},
		reports: make(ReportList),
	}
}

func (s *Service) GetOrNew(name string) *Report {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r, ok := s.reports[name]; ok {
		return r
	}

	s.reports[name] = &Report{}

	return s.reports[name]
}

func (s *Service) Update(key string, report *Report) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if report == nil {
		return fmt.Errorf("report param cannot be nil")
	}

	if r, ok := s.reports[key]; ok {
		r.Fixed += report.Fixed
		r.Errors += report.Errors
		r.Warnings += report.Warnings
		r.TotalFiles += report.TotalFiles
	}

	s.reports[key] = report

	return nil
}

func (s *Service) GetReports() ReportList {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.reports
}

type ReportList map[string]*Report

type Report struct {
	TotalFiles uint
	Errors     uint
	Warnings   uint
	Fixed      uint
}
