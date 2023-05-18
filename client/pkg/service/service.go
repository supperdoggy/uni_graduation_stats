package service

import (
	"github.com/supperdoggy/diploma_university_statistics_tool/client/pkg/clients/unistats"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

type IService interface {
	Schools() (*rest.ListSchoolsResponse, error)
	TopCompanies(school string) (*rest.ListSchoolsTopCompaniesResponse, error)
	TopHiredDegrees(school, company string) (*rest.TopHiredDegreesResponse, error)
}

type service struct {
	api unistats.IUniStats
	log *zap.Logger
}

func NewService(log *zap.Logger, api unistats.IUniStats) *service {
	return &service{
		api: api,
		log: log,
	}
}

func (s *service) Schools() (*rest.ListSchoolsResponse, error) {
	schools, err := s.api.SchoolList()
	if err != nil {
		s.log.Error("Error while getting schools", zap.Error(err))
		return nil, err
	}

	return schools, nil
}

func (s *service) TopCompanies(school string) (*rest.ListSchoolsTopCompaniesResponse, error) {
	companies, err := s.api.TopCompanies(school)
	if err != nil {
		s.log.Error("Error while getting companies", zap.Error(err))
		return nil, err
	}

	return companies, nil
}

func (s *service) TopHiredDegrees(school, company string) (*rest.TopHiredDegreesResponse, error) {
	degrees, err := s.api.TopHiredDegrees(school, company)
	if err != nil {
		s.log.Error("Error while getting degrees", zap.Error(err))
		return nil, err
	}

	return degrees, nil
}
