package storage

import (
	"context"
	"sync"

	"github.com/supperdoggy/diploma_university_statistics_tool/models/db"
	"github.com/supperdoggy/diploma_university_statistics_tool/models/rest"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type IMongoDB interface {

	// Users
	CreateUser(ctx context.Context, email, fullname string, password []byte) (*db.User, error)
	DeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, id, email string, password []byte) error
	GetUser(ctx context.Context, id string) (*db.User, error)
	GetUserByEmail(ctx context.Context, email string) (*db.User, error)

	// auth
	NewToken(ctx context.Context, userID string) (string, error)
	CheckToken(ctx context.Context, token string) (bool, string)

	// email codes
	NewEmailCode(ctx context.Context, email, code string) error
	CheckEmailCode(ctx context.Context, email, code string) (bool, error)

	// Education
	ListSchools(ctx context.Context) ([]rest.ListSchools, error)
	ListSchoolsTopCompanies(ctx context.Context, school string) ([]rest.ListSchoolsTopCompanies, error)
	ListJobsBySchool(ctx context.Context, school string) ([]rest.ListJobsBySchool, error)
	CorrelationBetweenDegreeAndTitle(ctx context.Context, school string) ([]rest.CorrelationDegreeAndTitle, error)
	SchoolDegrees(ctx context.Context, school string) ([]rest.SchoolDegrees, error)

	// Companies
	ListCompanies(ctx context.Context) ([]rest.ListCompanies, error)
	ListCompaniesTopSchools(ctx context.Context, company string) ([]rest.ListCompaniesTopSchools, error)
	TopHiredDegreesByCompany(ctx context.Context, company, school string) ([]rest.TopHiredDegrees, error)
}

type obj map[string]interface{}

type tokenCache struct {
	m   map[string]db.Token
	mut sync.Mutex
}

type emailCodeCache struct {
	m   map[string]db.EmailCode
	mut sync.Mutex
}

type mongodb struct {
	// The database connection string
	connStr string

	// The mongo client
	client *mongo.Client

	// collections
	students *mongo.Collection
	users    *mongo.Collection

	cache          tokenCache
	emailCodeCache emailCodeCache

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
		connStr:  connStr,
		client:   client,
		log:      log,
		users:    client.Database("stud").Collection("creds"),
		students: client.Database("stud").Collection("users"),

		cache: tokenCache{m: make(map[string]db.Token)},
	}, nil
}
