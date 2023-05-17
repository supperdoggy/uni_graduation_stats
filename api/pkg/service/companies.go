package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
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
	companies, err := s.db.ListCompaniesTopSchools(ctx, company)
	if err != nil {
		s.log.Error("error while getting companies", zap.Error(err))
		return nil, err
	}

	return companies, nil
}
