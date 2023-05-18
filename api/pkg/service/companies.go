package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.uber.org/zap"
)

func (s *service) ListCompanies(ctx context.Context) ([]rest.ListCompanies, error) {
	companies, err := s.db.ListCompanies(ctx)
	if err != nil {
		s.log.Error("error while getting companies", zap.Error(err))
		return nil, err
	}

	return companies, nil
}

func (s *service) ListCompaniesTopSchools(ctx context.Context, company string) ([]rest.ListCompaniesTopSchools, error) {
	schools, err := s.db.ListCompaniesTopSchools(ctx, company)
	if err != nil {
		s.log.Error("error while getting schools", zap.Error(err))
		return nil, err
	}

	return schools, nil
}

func (s *service) TopHiredDegreesByCompany(ctx context.Context, company, school string) ([]rest.TopHiredDegrees, error) {
	degrees, err := s.db.TopHiredDegreesByCompany(ctx, company, school)
	if err != nil {
		s.log.Error("error while getting degrees", zap.Error(err))
		return nil, err
	}

	return degrees, nil
}
