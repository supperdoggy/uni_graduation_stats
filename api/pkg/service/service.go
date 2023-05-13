package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/storage"
	"go.uber.org/zap"
)

type IService interface {
	// Universities
	ListSchools(ctx context.Context) ([]rest.ListUniversitiesSchools, error)

	// Companies
	ListCompanies(ctx context.Context) ([]rest.ListCompanies, error)
	ListCompaniesTopUniversities(ctx context.Context, company string) ([]rest.ListCompaniesTopUniversities, error)
}

type service struct {
	// storage interface
	db storage.IMongoDB

	// logger
	log *zap.Logger
}

func NewService(db storage.IMongoDB, log *zap.Logger) IService {
	return &service{
		db:  db,
		log: log,
	}
}
