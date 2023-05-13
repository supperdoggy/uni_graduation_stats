package storage

import (
	"context"

	"github.com/supperdoggy/diploma_university_statistics_tool/api/pkg/models/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type IMongoDB interface {
	// Education
	ListSchools(ctx context.Context) ([]rest.ListUniversitiesSchools, error)

	// Companies
	ListCompanies(ctx context.Context) ([]rest.ListCompanies, error)
	ListCompaniesTopUniversities(ctx context.Context, company string) ([]rest.ListCompaniesTopUniversities, error)
}

type mongodb struct {
	// The database connection string
	connStr string

	// The mongo client
	client *mongo.Client

	// collections
	users *mongo.Collection

	// The logger
	log *zap.Logger
}

func NewMongoDB(ctx context.Context, connStr string, log *zap.Logger) (IMongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))
	if err != nil {
		log.Error("error connecting to mongodb", zap.Error(err))
		return nil, err
	}

	return &mongodb{
		connStr: connStr,
		client:  client,
		log:     log,
		users:   client.Database("stud").Collection("users"),
	}, nil
}
