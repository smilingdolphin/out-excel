package services

import (
	"github.com/fpay/gopress"
)

const (
	// ExcelServiceName is the identity of excel service
	ExcelServiceName = "excel"
)

// ExcelService type
type ExcelService struct {
	// Uncomment this line if this service has dependence on other services in the container
	// c *gopress.Container
}

// NewExcelService returns instance of excel service
func NewExcelService() *ExcelService {
	return new(ExcelService)
}

// ServiceName is used to implements gopress.Service
func (s *ExcelService) ServiceName() string {
	return ExcelServiceName
}

// RegisterContainer is used to implements gopress.Service
func (s *ExcelService) RegisterContainer(c *gopress.Container) {
	// Uncomment this line if this service has dependence on other services in the container
	// s.c = c
}

func (s *ExcelService) SampleMethod() string {
	return "Excel Service"
}
