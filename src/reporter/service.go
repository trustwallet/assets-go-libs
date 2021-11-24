package reporter

import (
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
