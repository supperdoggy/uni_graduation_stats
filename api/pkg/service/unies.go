package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
)

func (s *service) ListSchools(ctx context.Context) ([]rest.ListUniversitiesSchools, error) {
	return s.db.ListSchools(ctx)
}
