package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.uber.org/zap"
)

func (s *service) ListSchools(ctx context.Context) ([]rest.ListSchools, error) {
	l, err := s.db.ListSchools(ctx)
	if err != nil {
		s.log.Error("error while getting schools", zap.Error(err))
		return nil, err
	}

	return l, nil
}

func (s *service) ListSchoolsTopCompanies(ctx context.Context, school string) ([]rest.ListSchoolsTopCompanies, error) {
	l, err := s.db.ListSchoolsTopCompanies(ctx, school)
	if err != nil {
		s.log.Error("error while getting schools", zap.Error(err))
		return nil, err
	}

	return l, nil
}
