package service

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/storage"
	"go.uber.org/zap"
)

type IService interface {
	ListSchools(ctx context.Context) ([]rest.ListUniversitiesSchools, error)
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
